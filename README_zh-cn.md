# typescript-go-internal

[English](README.md) | 简体中文

[![Test](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml/badge.svg)](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml)
[![codecov](https://codecov.io/gh/buke/typescript-go-internal/graph/badge.svg)](https://codecov.io/gh/buke/typescript-go-internal)
[![Go Report Card](https://goreportcard.com/badge/github.com/buke/typescript-go-internal)](https://goreportcard.com/report/github.com/buke/typescript-go-internal)
[![Go Reference](https://pkg.go.dev/badge/github.com/buke/typescript-go-internal.svg)](https://pkg.go.dev/github.com/buke/typescript-go-internal)

将 `microsoft/typescript-go` 的内部 Go 包以稳定的导入路径对外暴露，使外部模块可以依赖它们。

## 概述

本仓库将 `microsoft/typescript-go/internal` 镜像并适配到 `pkg/...` 目录，使这些包可以通过 `github.com/buke/typescript-go-internal/pkg/...` 导入。它与上游保持紧密同步，同时保持仓库自包含，便于 CI 和外部使用。

注意事项：
- 本项目独立运作，与 Microsoft 无关联。
- API 接口仍在演进中（v0 语义）：在跟踪上游变化时可能发生破坏性更改。

## 仓库内容

- `pkg/` — 镜像的内部包，已公开并可导入。
- `testdata/` — 从 `microsoft/typescript-go/testdata` 复制的上游测试数据和基线文件，用于可复现的测试。
- `scripts/sync-internal.sh` — 同步脚本，用于复制源码、重写导入路径、规范化 `//go:generate` 指令并同步测试数据。
- `.github/workflows/` — CI 工作流，包括测试和覆盖率上传。

## 从上游同步（维护者）

同步脚本执行以下操作：
- 复制 `microsoft/typescript-go/internal` → `pkg`
- 重写导入路径：`.../internal/...` → `.../pkg/...`
- 规范化 `//go:generate` 指令（使用 `go run <module>@latest`）
- 复制 `microsoft/typescript-go/testdata` → `testdata`
- 在生成前后运行 `go mod tidy`

命令：
```bash
./scripts/sync-internal.sh
```

同步后正常提交更改即可。基线文件位于 `testdata/baselines` 目录。

## 持续集成

- GitHub Actions 在推送和 PR 时运行 `go test` 并生成覆盖率报告。
- 覆盖率上传至 Codecov：https://codecov.io/gh/buke/typescript-go-internal

## 许可证

本项目采用 Apache License 2.0 许可 - 详见 [LICENSE](./LICENSE) 文件。

### 归属声明

本仓库包含基于以下项目的衍生作品：
- [`microsoft/typescript-go`](https://github.com/microsoft/typescript-go) (Apache 2.0)  
  Copyright (c) Microsoft Corporation
- [`microsoft/TypeScript`](https://github.com/microsoft/TypeScript) (Apache 2.0)  
  Copyright (c) Microsoft Corporation

完整归属详情请参见 [NOTICE](./NOTICE)。

**本项目与 Microsoft Corporation 无关联或背书关系。**