package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestNavto_excludeLib3(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: /index.ts
function [|parseInt|](s: string): number {}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "parseInt",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:     "parseInt",
					Kind:     lsproto.SymbolKindFunction,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		},
	})
}
