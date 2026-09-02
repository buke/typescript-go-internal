package build

import (
	"github.com/buke/typescript-go-internal/v7/pkg/ast"
	"github.com/buke/typescript-go-internal/v7/pkg/compiler"
	"github.com/buke/typescript-go-internal/v7/pkg/diagnostics"
	"github.com/buke/typescript-go-internal/v7/pkg/tsoptions"
	"github.com/buke/typescript-go-internal/v7/pkg/tspath"
	"github.com/buke/typescript-go-internal/v7/pkg/vfs"
)

type compilerHost struct {
	host  *host
	trace func(msg *diagnostics.Message, args ...any)
}

var _ compiler.CompilerHost = (*compilerHost)(nil)

func (h *compilerHost) FS() vfs.FS {
	return h.host.FS()
}

func (h *compilerHost) DefaultLibraryPath() string {
	return h.host.DefaultLibraryPath()
}

func (h *compilerHost) GetCurrentDirectory() string {
	return h.host.GetCurrentDirectory()
}

func (h *compilerHost) Trace(msg *diagnostics.Message, args ...any) {
	h.trace(msg, args...)
}

func (h *compilerHost) GetSourceFile(opts ast.SourceFileParseOptions) *ast.SourceFile {
	return h.host.GetSourceFile(opts)
}

func (h *compilerHost) GetResolvedProjectReference(fileName string, path tspath.Path) *tsoptions.ParsedCommandLine {
	return h.host.GetResolvedProjectReference(fileName, path)
}
