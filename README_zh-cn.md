# typescript-go-internal

[English](README.md) | 简体中文

[![Test](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml/badge.svg)](https://github.com/buke/typescript-go-internal/actions/workflows/go-test.yml)
[![codecov](https://codecov.io/gh/buke/typescript-go-internal/graph/badge.svg)](https://codecov.io/gh/buke/typescript-go-internal)
[![Go Report Card](https://goreportcard.com/badge/github.com/buke/typescript-go-internal)](https://goreportcard.com/report/github.com/buke/typescript-go-internal)
[![Go Reference](https://pkg.go.dev/badge/github.com/buke/typescript-go-internal/v7.svg)](https://pkg.go.dev/github.com/buke/typescript-go-internal/v7)

将 `microsoft/TypeScript`（`tsc/internal`）的内部 Go 包以稳定的导入路径对外暴露，使外部模块可以依赖它们。

## 概述

本仓库将 `microsoft/TypeScript/tsc/internal` 镜像并适配到 `pkg/...`，使这些包可以通过 `github.com/buke/typescript-go-internal/v7/pkg/...` 导入。它跟随 TypeScript **7.x 发版 tag**（当前为 `v7.0.2`），并保持仓库自包含，便于 CI 和外部使用。

注意事项：
- 本项目独立运作，与 Microsoft 无关联。
- Go 模块主版本与 TypeScript 主版本对齐（`/v7` ↔ TypeScript 7.x）。
- 迁移前的导入路径（`github.com/buke/typescript-go-internal/pkg/...`）仅存在于历史 tag。

## 安装

```bash
go get github.com/buke/typescript-go-internal/v7@v7.0.2
```

```go
import "github.com/buke/typescript-go-internal/v7/pkg/ast"
```

## 仓库内容

- `pkg/` — 镜像的内部包，已公开并可导入。
- `testdata/` — 从 `microsoft/TypeScript/tsc/testdata` 复制的上游测试数据与基线。
- `scripts/sync-internal.sh` — 同步脚本，用于复制源码、重写导入路径、规范化 `//go:generate` 并同步测试数据。
- `.github/workflows/` — CI、上游同步、发版打 tag 与 GoReleaser。

## 从上游同步（维护者）

将 `microsoft/TypeScript` 钉到发版 tag 后执行：

```bash
./scripts/sync-internal.sh
```

同步脚本会：

- 复制 `microsoft/TypeScript/tsc/internal` → `pkg`
- 将 `github.com/microsoft/typescript-go/internal` 或 `github.com/microsoft/TypeScript/tsc/internal` 重写为 `<module>/pkg`
- 规范化 `//go:generate` 指令
- 复制 `microsoft/TypeScript/tsc/testdata` → `testdata`
- 在生成前后运行 `go mod tidy`
- 运行 `scripts/post-sync.sh`，重新应用本地 harness 补丁（如稳定的 help 版本 padding）并刷新与版本相关的基线
- 设置 `SYNC_POST=0` 可跳过 post-sync

自动化：**Sync** workflow 会检查新的 TypeScript ≥ 7.0.0 稳定 tag 并开带 `sync` label 的 PR。合入后 **Release Tag** 会推送同名 git tag，**GoReleaser** 会创建 GitHub Release（仅库，无二进制）。

## 持续集成

- GitHub Actions 在推送和 PR 时运行 `go test` 并生成覆盖率报告。
- 覆盖率上传至 Codecov：https://codecov.io/gh/buke/typescript-go-internal

## 许可证

本项目采用 Apache License 2.0 许可 - 详见 [LICENSE](./LICENSE) 文件。

### 归属声明

本仓库包含基于以下项目的衍生作品：
- [`microsoft/TypeScript`](https://github.com/microsoft/TypeScript) (Apache 2.0)  
  Copyright (c) Microsoft Corporation
- 历史上以 [`microsoft/typescript-go`](https://github.com/microsoft/typescript-go) 发布的暂存树 (Apache 2.0)  
  Copyright (c) Microsoft Corporation

完整归属详情请参见 [NOTICE](./NOTICE)。

**本项目与 Microsoft Corporation 无关联或背书关系。**
