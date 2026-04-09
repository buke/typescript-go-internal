package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNavto_excludeLib1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: /index.ts
import { weirdName as otherName } from "bar";
const [|weirdName: number = 1|];
// @filename: /tsconfig.json
{}
// @filename: /node_modules/bar/index.d.ts
export const [|weirdName: number|];
// @filename: /node_modules/bar/package.json
{}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "weirdName",
			Preferences: &lsutil.UserPreferences{ExcludeLibrarySymbolsInNavTo: core.TSFalse},
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:     "weirdName",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[0].LSLocation(),
				},
				{
					Name:     "weirdName",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[1].LSLocation(),
				},
			}),
		},
	})
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "weirdName",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:     "weirdName",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		},
	})
}
