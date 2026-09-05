#!/usr/bin/env bash
# Post-sync fixes applied after mirroring microsoft/TypeScript/tsc into this repo.
# Kept separate so upstream copies do not wipe local harness stability patches.
#
# Current fixes:
#   1) Re-apply scripts/patches/*.patch (TS_TEST_VERSION help padding)
#   2) Refresh version-sensitive tsc help baselines to match FakeTSVersion padding
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

PATCH_DIR="${ROOT_DIR}/scripts/patches"
POST_SYNC="${SYNC_POST:-1}"

if [[ "${POST_SYNC}" != "1" ]]; then
  echo "Skip post-sync (SYNC_POST=${POST_SYNC})"
  exit 0
fi

echo "Running post-sync fixes..."

apply_patches() {
  if [[ ! -d "${PATCH_DIR}" ]]; then
    return 0
  fi
  local patch
  shopt -s nullglob
  for patch in "${PATCH_DIR}"/*.patch; do
    echo "Applying $(basename "$patch")..."
    if git apply --check "$patch" >/dev/null 2>&1; then
      git apply "$patch"
    elif grep -q 'TS_TEST_VERSION' pkg/execute/tsc/help.go 2>/dev/null; then
      echo "  already applied; skipping"
    else
      echo "  ERROR: failed to apply $patch" >&2
      git apply --check "$patch" || true
      exit 1
    fi
  done
}

refresh_help_baselines() {
  echo "Refreshing version-sensitive help baselines..."
  # Failure writes testdata/baselines/local/...; we then accept into reference/.
  GOWORK=off go test ./pkg/execute/tsctests/ -count=1 \
    -run 'TestTscCommandline$/tsc/show_help_with_ExitStatus' >/dev/null 2>&1 || true

  local loc="testdata/baselines/local/tsc/commandLine"
  local ref="testdata/baselines/reference/tsc/commandLine"
  mkdir -p "$ref"
  local f
  for f in \
    show-help-with-ExitStatus.DiagnosticsPresent_OutputsSkipped.js \
    show-help-with-ExitStatus.DiagnosticsPresent_OutputsSkipped-when-host-cannot-provide-terminal-width.js
  do
    if [[ -f "${loc}/${f}" ]]; then
      cp "${loc}/${f}" "${ref}/${f}"
      echo "  accepted ${f}"
    fi
  done

  GOWORK=off go test ./pkg/execute/tsctests/ -count=1 \
    -run 'TestTscCommandline$/tsc/show_help_with_ExitStatus'
}

apply_patches
refresh_help_baselines

echo "Post-sync complete."
