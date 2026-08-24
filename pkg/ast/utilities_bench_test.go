package ast_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/ast"
	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/parser"
	"github.com/buke/typescript-go-internal/pkg/testutil/fixtures"
	"github.com/buke/typescript-go-internal/pkg/tspath"
	"github.com/buke/typescript-go-internal/pkg/vfs/osvfs"
)

func BenchmarkGetCombinedFlags(b *testing.B) {
	for _, f := range fixtures.BenchFixtures {
		b.Run(f.Name(), func(b *testing.B) {
			f.SkipIfNotExist(b)

			fileName := tspath.GetNormalizedAbsolutePath(f.Path(), "/")
			path := tspath.ToPath(fileName, "/", osvfs.FS().UseCaseSensitiveFileNames())
			sourceText := f.ReadFile(b)
			scriptKind := core.GetScriptKindFromFileName(fileName)

			sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
				FileName: fileName,
				Path:     path,
			}, sourceText, scriptKind)

			var decls []*ast.Node
			var collect ast.Visitor
			collect = func(n *ast.Node) bool {
				if ast.IsDeclaration(n) {
					decls = append(decls, n)
				}
				n.ForEachChild(collect)
				return false
			}
			sourceFile.AsNode().ForEachChild(collect)

			for b.Loop() {
				for _, n := range decls {
					_ = ast.GetCombinedNodeFlags(n)
					_ = ast.GetCombinedModifierFlags(n)
				}
			}
		})
	}
}
