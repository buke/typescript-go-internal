package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNavigateItemsLet(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noLib: true
let [|c|] = 10;
function foo() {
    let [|d|] = 10;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "c",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:     "c",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		}, {
			Pattern:     "d",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:          "d",
					Kind:          lsproto.SymbolKindVariable,
					Location:      f.Ranges()[1].LSLocation(),
					ContainerName: new("foo"),
				},
			}),
		},
	})
}
