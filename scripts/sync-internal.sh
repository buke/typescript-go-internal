#!/usr/bin/env bash
# Sync microsoft/typescript-go/internal -> pkg, rewrite imports, and normalize go:generate directives.
# Generation modes:
#   SYNC_GENERATE=full  (default) run `go generate ./pkg/...` and create a temporary _submodules symlink
#   SYNC_GENERATE=light skip packages pkg/bundled and pkg/diagnostics
#   SYNC_GENERATE=skip  skip generation

set -euo pipefail
[[ "${SYNC_VERBOSE:-0}" == "1" ]] && set -x

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

UPSTREAM_SUBMODULE="microsoft/typescript-go"
SRC_DIR="${UPSTREAM_SUBMODULE}/internal"
DEST_DIR="pkg"
GENERATE_MODE="${SYNC_GENERATE:-full}"
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

# Read module path
if [[ ! -f "go.mod" ]]; then
  echo "go.mod not found. Please run: go mod init <module>"
  exit 1
fi
MODULE_PATH="$(awk '/^module /{print $2}' go.mod)"
if [[ -z "$MODULE_PATH" ]]; then
  echo "Failed to read module path from go.mod"
  exit 1
fi

# Submodule check (optional)
# if [[ ! -d "$UPSTREAM_SUBMODULE/.git" ]]; then
#   echo "Submodule not initialized: ${UPSTREAM_SUBMODULE}"
#   echo "Run: git submodule update --init --recursive"
#   exit 1
# fi

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

echo "Sync complete."