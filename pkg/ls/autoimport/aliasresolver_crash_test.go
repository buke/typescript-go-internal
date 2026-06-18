package autoimport

import (
	"context"
	"testing"

	"github.com/buke/typescript-go-internal/pkg/ast"
	"github.com/buke/typescript-go-internal/pkg/binder"
	"github.com/buke/typescript-go-internal/pkg/checker"
	"github.com/buke/typescript-go-internal/pkg/compiler"
	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/module"
	"github.com/buke/typescript-go-internal/pkg/packagejson"
	"github.com/buke/typescript-go-internal/pkg/parser"
	"github.com/buke/typescript-go-internal/pkg/tspath"
	"github.com/buke/typescript-go-internal/pkg/vfs"
	"github.com/buke/typescript-go-internal/pkg/vfs/vfstest"
)

type fakeCloneHost struct {
	fs vfs.FS
}

func (h *fakeCloneHost) FS() vfs.FS                  { return h.fs }
func (h *fakeCloneHost) GetCurrentDirectory() string { return "/" }
func (h *fakeCloneHost) GetDefaultProject(path tspath.Path) (tspath.Path, *compiler.Program) {
	return "", nil
}

func (h *fakeCloneHost) GetProgramForProject(projectPath tspath.Path) *compiler.Program { return nil }

func (h *fakeCloneHost) GetPackageJson(fileName string) *packagejson.InfoCacheEntry { return nil }

func (h *fakeCloneHost) GetSourceFile(fileName string, path tspath.Path) *ast.SourceFile { return nil }
func (h *fakeCloneHost) Dispose()                                                        {}

var _ RegistryCloneHost = (*fakeCloneHost)(nil)

// Regression test for microsoft/typescript-go#4322.
//
// During auto-import export extraction, the checker is built on top of an
// aliasResolver standing in for a real program. This file has a type error, and
// extracting exports should still complete without crashing.
func TestAliasResolverGetDiagnosticsDoesNotPanic(t *testing.T) {
	t.Parallel()

	const fileName = "/pkg/index.ts"
	text := "declare function f(arg: { a: string }): () => void;\nexport const x = f({ a: 1 });\n"

	fs := vfstest.FromMap(map[string]string{fileName: text}, true /*useCaseSensitiveFileNames*/)
	host := &fakeCloneHost{fs: fs}

	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: fileName,
		Path:     tspath.Path(fileName),
	}, text, core.ScriptKindTS)
	binder.BindSourceFile(sourceFile)

	resolver := module.NewResolver(host, core.EmptyCompilerOptions, "", "")
	r := newAliasResolver(
		[]*ast.SourceFile{sourceFile},
		nil,
		host,
		resolver,
		func(f string) tspath.Path { return tspath.Path(f) },
		func(ast.HasFileName, string) {},
	)

	ch, _ := checker.NewChecker(r, nil)

	// Type-checking this file's diagnostics must not panic.
	ch.GetDiagnostics(context.Background(), sourceFile)
}
