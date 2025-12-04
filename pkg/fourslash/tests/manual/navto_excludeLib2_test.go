package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNavto_excludeLib2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: /index.ts
import { [|someName as weirdName|] } from "bar";
// @filename: /tsconfig.json
{}
// @filename: /node_modules/bar/index.d.ts
export const someName: number;
// @filename: /node_modules/bar/package.json
{}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "weirdName",
			Preferences: &lsutil.UserPreferences{ExcludeLibrarySymbolsInNavTo: false},
			Exact: PtrTo([]*lsproto.SymbolInformation{
				{
					Name:     "weirdName",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		},
	})
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "weirdName",
			Preferences: nil,
			Exact: PtrTo([]*lsproto.SymbolInformation{
				{
					Name:     "weirdName",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		},
	})
}
