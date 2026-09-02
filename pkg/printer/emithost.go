package printer

import (
	"github.com/buke/typescript-go-internal/v7/pkg/ast"
	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/tsoptions"
	"github.com/buke/typescript-go-internal/v7/pkg/tspath"
)

// NOTE: EmitHost operations must be thread-safe
type EmitHost interface {
	Options() *core.CompilerOptions
	SourceFiles() []*ast.SourceFile
	UseCaseSensitiveFileNames() bool
	GetCurrentDirectory() string
	CommonSourceDirectory() string
	IsEmitBlocked(file string) bool
	WriteFile(fileName string, text string) error
	GetEmitModuleFormatOfFile(file ast.HasFileName) core.ModuleKind
	GetEmitResolver() EmitResolver
	GetProjectReferenceFromSource(path tspath.Path) *tsoptions.SourceOutputAndProjectReference
	IsSourceFileFromExternalLibrary(file *ast.SourceFile) bool
}
