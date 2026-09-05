# typescript-go-internal

English | [简体中文](README_zh-cn.md)

[![Test](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml/badge.svg)](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml)
[![codecov](https://codecov.io/gh/buke/typescript-go-internal/graph/badge.svg)](https://codecov.io/gh/buke/typescript-go-internal)
[![Go Report Card](https://goreportcard.com/badge/github.com/buke/typescript-go-internal)](https://goreportcard.com/report/github.com/buke/typescript-go-internal)
[![Go Reference](https://pkg.go.dev/badge/github.com/buke/typescript-go-internal/v7.svg)](https://pkg.go.dev/github.com/buke/typescript-go-internal/v7)

Expose selected internal Go packages from `microsoft/TypeScript` (`tsc/internal`) under stable import paths so external modules can depend on them.

## Overview

This repository mirrors and adapts `microsoft/TypeScript/tsc/internal` into `pkg/...` to make those packages importable as `github.com/buke/typescript-go-internal/v7/pkg/...`. It tracks TypeScript **7.x release tags** (currently `v7.0.2`) and remains self-contained for CI and external use.

Notes:
- This project is independent and not affiliated with Microsoft.
- Go module major versions align with TypeScript majors (`/v7` ↔ TypeScript 7.x).
- Pre-`/v7` imports (`github.com/buke/typescript-go-internal/pkg/...`) remain available only on historical tags.

## Install

```bash
go get github.com/buke/typescript-go-internal/v7@v7.0.2
```

```go
import "github.com/buke/typescript-go-internal/v7/pkg/ast"
```

## What’s Inside

- `pkg/` — mirrored internal packages made public and importable.
- `testdata/` — upstream fixtures and baselines copied from `microsoft/TypeScript/tsc/testdata`.
- `scripts/sync-internal.sh` — sync script to copy sources, rewrite imports, normalize `//go:generate`, and bring testdata.
- `.github/workflows/` — CI, upstream sync, release tagging, and GoReleaser.

## Syncing From Upstream (maintainers)

Pin `microsoft/TypeScript` to a release tag, then run:

```bash
./scripts/sync-internal.sh
```

The sync script:

- Copies `microsoft/TypeScript/tsc/internal` → `pkg`
- Rewrites imports from either `github.com/microsoft/typescript-go/internal` or `github.com/microsoft/TypeScript/tsc/internal` → `<module>/pkg`
- Normalizes `//go:generate` directives
- Copies `microsoft/TypeScript/tsc/testdata` → `testdata`
- Runs `go mod tidy` pre/post generation
- Runs `scripts/post-sync.sh` to re-apply local harness patches (e.g. stable help-version padding) and refresh version-sensitive baselines
- Set `SYNC_POST=0` to skip post-sync

Automated sync: the **Sync** workflow checks for new stable TypeScript ≥ 7.0.0 tags and opens a PR labeled `sync`. After merge, **Release Tag** pushes a matching git tag and **GoReleaser** publishes a GitHub Release (library-only; no binaries).

## Continuous Integration

- GitHub Actions runs `go test` with coverage on pushes and PRs.
- Coverage is uploaded to Codecov: https://codecov.io/gh/buke/typescript-go-internal

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](./LICENSE) file for details.

### Attribution

This repository contains derivative works based on:
- [`microsoft/TypeScript`](https://github.com/microsoft/TypeScript) (Apache 2.0)  
  Copyright (c) Microsoft Corporation
- Historical staging tree formerly published as [`microsoft/typescript-go`](https://github.com/microsoft/typescript-go) (Apache 2.0)  
  Copyright (c) Microsoft Corporation

See [NOTICE](./NOTICE) for full attribution details.

**This project is not affiliated with or endorsed by Microsoft Corporation.**
