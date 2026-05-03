#!/usr/bin/env bash
# Sync microsoft/typescript-go/internal -> pkg, rewrite imports, and normalize go:generate directives.
# Generation modes:
#   SYNC_GENERATE=full  (default) run `go generate ./pkg/...` and create a temporary _submodules symlink
#   SYNC_GENERATE=light skip packages pkg/bundled and pkg/diagnostics
#   SYNC_GENERATE=skip  skip generation
# Module alignment modes:
#   SYNC_MOD_ALIGN=shared (default) align versions of direct dependencies shared by root and upstream go.mod
#   SYNC_MOD_ALIGN=strict mirror upstream direct dependencies into root go.mod (add/update/drop)
#   SYNC_MOD_ALIGN=off    skip go.mod dependency alignment
# Drift check:
#   SYNC_MOD_DRIFT_CHECK=1 fail if shared direct dependencies still drift after sync

set -euo pipefail
[[ "${SYNC_VERBOSE:-0}" == "1" ]] && set -x

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

UPSTREAM_SUBMODULE="microsoft/typescript-go"
UPSTREAM_MOD_FILE="${UPSTREAM_SUBMODULE}/go.mod"
SRC_DIR="${UPSTREAM_SUBMODULE}/internal"
DEST_DIR="pkg"
GENERATE_MODE="${SYNC_GENERATE:-full}"
MOD_ALIGN_MODE="${SYNC_MOD_ALIGN:-shared}"
MOD_DRIFT_CHECK="${SYNC_MOD_DRIFT_CHECK:-0}"
TS_SYMLINK="${SYNC_TS_SYMLINK:-1}"
TS_SYMLINK_CREATED=0
TS_SYMLINK_EXISTED=0
TS_SYMLINK_POINTS_TARGET=0
TS_SYMLINK_CLEAN_EXISTING="${SYNC_TS_SYMLINK_CLEAN:-1}"

need_cmd() { command -v "$1" >/dev/null 2>&1 || { echo "Missing command: $1"; exit 1; }; }
need_cmd git
need_cmd rsync
need_cmd go
need_cmd awk
need_cmd sed
need_cmd grep
need_cmd xargs
need_cmd find
need_cmd ln
need_cmd dirname
need_cmd basename
need_cmd sort
need_cmd mktemp

# Read module path
if [[ ! -f "go.mod" ]]; then
  echo "go.mod not found. Please run: go mod init <module>"
  exit 1
fi
if [[ ! -f "${UPSTREAM_MOD_FILE}" ]]; then
  echo "Upstream go.mod not found: ${UPSTREAM_MOD_FILE}"
  exit 1
fi
MODULE_PATH="$(awk '/^module /{print $2}' go.mod)"
if [[ -z "$MODULE_PATH" ]]; then
  echo "Failed to read module path from go.mod"
  exit 1
fi

case "${MOD_ALIGN_MODE}" in
  off|shared|strict) ;;
  *)
    echo "Invalid SYNC_MOD_ALIGN value: ${MOD_ALIGN_MODE}. Expected: off|shared|strict"
    exit 1
    ;;
esac

case "${MOD_DRIFT_CHECK}" in
  0|1) ;;
  *)
    echo "Invalid SYNC_MOD_DRIFT_CHECK value: ${MOD_DRIFT_CHECK}. Expected: 0|1"
    exit 1
    ;;
esac

echo "Module: ${MODULE_PATH}"
echo "Sync:   ${SRC_DIR}  ->  ${DEST_DIR}"

# Wipe and copy internal -> pkg
rm -rf "${DEST_DIR}"
mkdir -p "${DEST_DIR}"
rsync -a --delete \
  --exclude '.git' \
  --exclude 'vendor' \
  "${SRC_DIR}/" "${DEST_DIR}/"

# sed -i compatibility
if [[ "$OSTYPE" == "darwin"* ]]; then
  SED_INPLACE=("sed" "-i" "")
  SED_EXT=(-E)
else
  SED_INPLACE=("sed" "-i")
  SED_EXT=(-r)
fi

# Rewrite imports: github.com/microsoft/typescript-go/internal/... -> <module>/pkg/...
echo "Rewriting imports..."
find "${DEST_DIR}" -type f -name "*.go" -print0 | xargs -0 "${SED_INPLACE[@]}" \
  -e "s#\"github.com/microsoft/typescript-go/internal#\"${MODULE_PATH}/pkg#g" \
  -e "s#'github.com/microsoft/typescript-go/internal#'${MODULE_PATH}/pkg#g" \
  -e "s#\`github.com/microsoft/typescript-go/internal#\`${MODULE_PATH}/pkg#g"

# Normalize go:generate: convert "go tool <module> ..." -> "go run <module>@latest ..."
# Only match module-like paths that contain a dot in the domain part.
echo "Fixing //go:generate directives (go tool -> go run ...@latest)..."

fix_generate_file() {
  local f="$1"
  if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS: sed -E -i ''
    sed -E -i '' \
      's#(//go:generate[[:space:]]+)go[[:space:]]+tool[[:space:]]+([A-Za-z0-9._-]+\.[A-Za-z0-9._-]+(/[A-Za-z0-9._@+-]+)+)([[:space:]]|$)#\1go run \2@latest\4#g' \
      "$f"
  else
    # Linux: sed -r -i
    sed -r -i \
      's#(//go:generate[[:space:]]+)go[[:space:]]+tool[[:space:]]+([A-Za-z0-9._-]+\.[A-Za-z0-9._-]+(/[A-Za-z0-9._@+-]+)+)([[:space:]]|$)#\1go run \2@latest\4#g' \
      "$f"
  fi
}

while IFS= read -r -d '' file; do
  fix_generate_file "$file"
done < <(find "${DEST_DIR}" -type f -name "*.go" -print0)

# moq can fail when type-checking transitive deps under newer Go syntax.
# Add -skip-ensure so generation only depends on interface parsing.
echo "Adjusting moq directives (-skip-ensure)..."
find "${DEST_DIR}" -type f -name "*.go" -print0 | xargs -0 "${SED_INPLACE[@]}" \
  -e "s#go run github.com/matryer/moq@latest #go run github.com/matryer/moq@latest -skip-ensure #g"

# Create symlink for generators that need TypeScript data files:
#   _submodules -> microsoft/typescript-go/_submodules
ensure_ts_symlink() {
  local target="${UPSTREAM_SUBMODULE}/_submodules"
  local link="_submodules"

  if [[ "${TS_SYMLINK}" != "1" ]]; then
    echo "Skip creating symlink (_submodules) due to SYNC_TS_SYMLINK=${TS_SYMLINK}"
    return 0
  fi
  if [[ ! -d "${target}" ]]; then
    echo "Skip symlink: target not found: ${target}"
    return 0
  fi

  if [[ -e "$link" || -L "$link" ]]; then
    TS_SYMLINK_EXISTED=1
    if [[ -L "$link" ]]; then
      local cur
      cur="$(readlink "$link" || true)"
      if [[ "$cur" == "$target" ]]; then
        TS_SYMLINK_POINTS_TARGET=1
        echo "Using existing ${link}; points to target."
      else
        echo "Using existing ${link}; no changes."
      fi
    else
      echo "Existing ${link} is not a symlink; no changes."
    fi
    return 0
  fi

  ln -s "$target" "$link"
  TS_SYMLINK_CREATED=1
  TS_SYMLINK_POINTS_TARGET=1
  echo "Created symlink: ${link} -> ${target}"
}

cleanup_ts_symlink() {
  # Removal conditions:
  #  1) The symlink was created by this script in this run; or
  #  2) It pre-existed, points to the expected target, and SYNC_TS_SYMLINK_CLEAN=1
  if [[ -L "_submodules" ]]; then
    local cur
    cur="$(readlink "_submodules" || true)"
    if [[ "${TS_SYMLINK_CREATED}" == "1" ]]; then
      rm -f "_submodules"
      echo "Removed symlink: _submodules (created-by-script)"
    elif [[ "${TS_SYMLINK_EXISTED}" == "1" && "${TS_SYMLINK_POINTS_TARGET}" == "1" && "${TS_SYMLINK_CLEAN_EXISTING}" == "1" && "$cur" == "${UPSTREAM_SUBMODULE}/_submodules" ]]; then
      rm -f "_submodules"
      echo "Removed symlink: _submodules (pre-existing but points-to-target)"
    fi
  fi
}

# Copy testdata from submodule into repo (self-contained for CI/tests)
copy_testdata() {
  local src="${UPSTREAM_SUBMODULE}/testdata"
  local dst="testdata"
  if [[ ! -d "${src}" ]]; then
    echo "Skip testdata copy: source not found: ${src}"
    return 0
  fi
  # If dest is a symlink, remove it
  if [[ -L "${dst}" ]]; then
    rm -f "${dst}"
  fi
  mkdir -p "${dst}"
  rsync -a --delete "${src}/" "${dst}/"
  echo "Copied testdata: ${src} -> ${dst}"
}

collect_direct_requires() {
  local mod_file="$1"
  awk '
    BEGIN { in_block=0 }
    /^[[:space:]]*require[[:space:]]*\(/ { in_block=1; next }
    in_block && /^[[:space:]]*\)/ { in_block=0; next }

    {
      raw=$0
      line=raw
      sub(/[[:space:]]*\/\/.*$/, "", line)
      gsub(/^[[:space:]]+|[[:space:]]+$/, "", line)
      if (line == "") next

      if (!in_block && line ~ /^require[[:space:]]+/ && line !~ /^require[[:space:]]*\(/) {
        if (raw ~ /\/\/[[:space:]]*indirect([[:space:]]|$)/) next
        sub(/^require[[:space:]]+/, "", line)
        n=split(line, a, /[[:space:]]+/)
        if (n >= 2) print a[1], a[2]
        next
      }

      if (in_block) {
        if (raw ~ /\/\/[[:space:]]*indirect([[:space:]]|$)/) next
        n=split(line, a, /[[:space:]]+/)
        if (n >= 2) print a[1], a[2]
      }
    }
  ' "$mod_file" | sort -k1,1
}

has_module() {
  local req_file="$1"
  local mod="$2"
  awk -v m="$mod" '$1==m { found=1 } END { exit found ? 0 : 1 }' "$req_file"
}

get_module_version() {
  local req_file="$1"
  local mod="$2"
  awk -v m="$mod" '$1==m { print $2; exit }' "$req_file"
}

align_go_mod_requires() {
  local mode="$1"
  local root_reqs
  local upstream_reqs
  local changed=0
  local mod
  local upstream_ver
  local root_ver

  root_reqs="$(mktemp)"
  upstream_reqs="$(mktemp)"

  collect_direct_requires "go.mod" > "$root_reqs"
  collect_direct_requires "${UPSTREAM_MOD_FILE}" > "$upstream_reqs"

  case "$mode" in
    off)
      echo "Skip go.mod dependency alignment (SYNC_MOD_ALIGN=off)"
      rm -f "$root_reqs" "$upstream_reqs"
      return 0
      ;;
    shared)
      echo "Aligning shared direct dependencies to upstream (SYNC_MOD_ALIGN=shared)..."
      while read -r mod upstream_ver; do
        [[ -z "${mod:-}" ]] && continue
        root_ver="$(get_module_version "$root_reqs" "$mod")"
        if [[ -n "$root_ver" && "$root_ver" != "$upstream_ver" ]]; then
          echo "  ${mod}: ${root_ver} -> ${upstream_ver}"
          GOWORK=off go mod edit -require="${mod}@${upstream_ver}"
          changed=1
        fi
      done < "$upstream_reqs"
      ;;
    strict)
      echo "Mirroring upstream direct dependencies (SYNC_MOD_ALIGN=strict)..."
      while read -r mod upstream_ver; do
        [[ -z "${mod:-}" ]] && continue
        root_ver="$(get_module_version "$root_reqs" "$mod")"
        if [[ "$root_ver" != "$upstream_ver" ]]; then
          if [[ -n "$root_ver" ]]; then
            echo "  ${mod}: ${root_ver} -> ${upstream_ver}"
          else
            echo "  add ${mod}@${upstream_ver}"
          fi
          GOWORK=off go mod edit -require="${mod}@${upstream_ver}"
          changed=1
        fi
      done < "$upstream_reqs"

      collect_direct_requires "go.mod" > "$root_reqs"
      while read -r mod root_ver; do
        [[ -z "${mod:-}" ]] && continue
        if ! has_module "$upstream_reqs" "$mod"; then
          echo "  drop ${mod} (not in upstream direct requires)"
          GOWORK=off go mod edit -droprequire="$mod" || true
          changed=1
        fi
      done < "$root_reqs"
      ;;
  esac

  if [[ "$changed" == "0" ]]; then
    echo "No go.mod direct dependency version updates were needed."
  fi

  rm -f "$root_reqs" "$upstream_reqs"
}

check_shared_mod_drift() {
  local root_reqs
  local upstream_reqs
  local drift_file
  local mod
  local root_ver
  local upstream_ver

  root_reqs="$(mktemp)"
  upstream_reqs="$(mktemp)"
  drift_file="$(mktemp)"

  collect_direct_requires "go.mod" > "$root_reqs"
  collect_direct_requires "${UPSTREAM_MOD_FILE}" > "$upstream_reqs"

  while read -r mod upstream_ver; do
    [[ -z "${mod:-}" ]] && continue
    root_ver="$(get_module_version "$root_reqs" "$mod")"
    if [[ -n "$root_ver" && "$root_ver" != "$upstream_ver" ]]; then
      printf "%s %s %s\n" "$mod" "$root_ver" "$upstream_ver" >> "$drift_file"
    fi
  done < "$upstream_reqs"

  if [[ -s "$drift_file" ]]; then
    echo "Shared direct dependency drift detected (root vs upstream):"
    while read -r mod root_ver upstream_ver; do
      echo "  - ${mod}: root=${root_ver}, upstream=${upstream_ver}"
    done < "$drift_file"
    rm -f "$root_reqs" "$upstream_reqs" "$drift_file"
    return 1
  fi

  echo "Shared direct dependency drift check passed."
  rm -f "$root_reqs" "$upstream_reqs" "$drift_file"
  return 0
}

# Align root go.mod against upstream go.mod before tidy/generate.
align_go_mod_requires "$MOD_ALIGN_MODE"

# Tidy before generation
echo "Running: go mod tidy (pre-generate)"
GOWORK="${GOWORK:-off}" go mod tidy

# Run go:generate according to mode
run_generate() {
  case "${GENERATE_MODE}" in
    skip)
      echo "Skip go:generate (SYNC_GENERATE=skip)"
      return 0
      ;;
    light)
      echo "Running light generate (skip pkg/bundled and pkg/diagnostics)"
      mapfile -t PKGS < <(go list ./pkg/... | grep -vE '/pkg/(bundled|diagnostics)(/|$)' || true)
      if [[ ${#PKGS[@]} -eq 0 ]]; then
        echo "No packages to generate (light mode)."
        return 0
      fi
      GOWORK=off go generate "${PKGS[@]}" || echo "Warning: failures occurred during light generate."
      ;;
    full|*)
      echo "Running full generate (with _submodules symlink)"
      ensure_ts_symlink
      trap cleanup_ts_symlink EXIT
      GOWORK=off go generate ./pkg/... || echo "Warning: failures occurred during full generate."
      cleanup_ts_symlink
      ;;
  esac
}

run_generate

# Tidy after generation
echo "Running: go mod tidy (post-generate)"
GOWORK=off go mod tidy || true

# Optional guardrail for CI: fail when shared direct dependencies still drift.
if [[ "$MOD_DRIFT_CHECK" == "1" ]]; then
  check_shared_mod_drift
fi

# Ensure testdata is present in repo (copy from submodule)
copy_testdata

echo "Sync complete."