# typescript-go-internal

English | [简体中文](README_zh-cn.md)

[![Test](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml/badge.svg)](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml)
[![codecov](https://codecov.io/gh/buke/typescript-go-internal/graph/badge.svg)](https://codecov.io/gh/buke/typescript-go-internal)
[![Go Report Card](https://goreportcard.com/badge/github.com/buke/typescript-go-internal)](https://goreportcard.com/report/github.com/buke/typescript-go-internal)
[![Go Reference](https://pkg.go.dev/badge/github.com/buke/typescript-go-internal.svg)](https://pkg.go.dev/github.com/buke/typescript-go-internal)

Expose selected internal Go packages from `microsoft/typescript-go` under stable import paths so external modules can depend on them.

## Overview

This repository mirrors and adapts `microsoft/typescript-go/internal` into `pkg/...` to make those packages importable as `github.com/buke/typescript-go-internal/pkg/...`. It keeps close parity with upstream while remaining self-contained for CI and external use.

Notes:
- This project is independent and not affiliated with Microsoft.
- The API surface is still evolving (v0 semantics): breaking changes may occur while tracking upstream.

## What’s Inside

- `pkg/` — mirrored internal packages made public and importable.
- `testdata/` — upstream fixtures and baselines copied from `microsoft/typescript-go/testdata` for reproducible tests.
- `scripts/sync-internal.sh` — sync script to copy sources, rewrite imports, normalize `//go:generate`, and bring testdata.
- `.github/workflows/` — CI workflows, including tests and coverage upload.

## Syncing From Upstream (maintainers)

The sync script performs:
- Copy `microsoft/typescript-go/internal` → `pkg`
- Rewrite imports from `.../internal/...` → `.../pkg/...`
- Normalize `//go:generate` directives (use `go run <module>@latest`)
- Copy `microsoft/typescript-go/testdata` → `testdata`
- Run `go mod tidy` pre/post generation

Command:
```bash
./scripts/sync-internal.sh
```

After syncing, commit changes normally. Baseline files live under `testdata/baselines`.

## Continuous Integration

- GitHub Actions runs `go test` with coverage on pushes and PRs.
- Coverage is uploaded to Codecov: https://codecov.io/gh/buke/typescript-go-internal

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](./LICENSE) file for details.

### Attribution

This repository contains derivative works based on:
- [`microsoft/typescript-go`](https://github.com/microsoft/typescript-go) (Apache 2.0)  
  Copyright (c) Microsoft Corporation
- [`microsoft/TypeScript`](https://github.com/microsoft/TypeScript) (Apache 2.0)  
  Copyright (c) Microsoft Corporation

See [NOTICE](./NOTICE) for full attribution details.

**This project is not affiliated with or endorsed by Microsoft Corporation.**