package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNavigateToSymbolIterator(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: es5
class C {
    [|[Symbol.iterator]() {}|]
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "iterator",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:          "iterator",
					Kind:          lsproto.SymbolKindMethod,
					Location:      f.Ranges()[0].LSLocation(),
					ContainerName: new("C"),
				},
			}),
		},
	})
}
