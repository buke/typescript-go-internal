# Sync patches

Patches in this directory are re-applied by `scripts/post-sync.sh` after every
`scripts/sync-internal.sh` run (unless `SYNC_POST=0`).

They encode local harness fixes that would otherwise be wiped when upstream
`tsc/internal` is copied into `pkg/`. Keep patches against a clean upstream-
synced tree (no local edits already applied).
