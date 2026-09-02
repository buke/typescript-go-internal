package binder

import (
	"runtime"
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/ast"
	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/parser"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil/fixtures"
	"github.com/buke/typescript-go-internal/v7/pkg/tspath"
	"github.com/buke/typescript-go-internal/v7/pkg/vfs/osvfs"
)

func BenchmarkBind(b *testing.B) {
	for _, f := range fixtures.BenchFixtures {
		b.Run(f.Name(), func(b *testing.B) {
			f.SkipIfNotExist(b)

			fileName := tspath.GetNormalizedAbsolutePath(f.Path(), "/")
			path := tspath.ToPath(fileName, "/", osvfs.FS().UseCaseSensitiveFileNames())
			sourceText := f.ReadFile(b)

			parseOptions := ast.SourceFileParseOptions{
				FileName: fileName,
				Path:     path,
			}
			scriptKind := core.GetScriptKindFromFileName(fileName)

			sourceFiles := make([]*ast.SourceFile, b.N)
			for i := range b.N {
				sourceFiles[i] = parser.ParseSourceFile(parseOptions, sourceText, scriptKind)
			}

			// The above parses do a lot of work; ensure GC is finished before we start collecting performance data.
			// GC must be called twice to allow things to settle.
			runtime.GC()
			runtime.GC()

			b.ResetTimer()
			for i := range b.N {
				BindSourceFile(sourceFiles[i])
			}
		})
	}
}
