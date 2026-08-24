package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportStatementCompletionUsesNamedImport(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
export interface I {}
// @Filename: 1.ts
import * as u from "./a";
[|import I/*a*/|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	preferences := lsutil.NewDefaultUserPreferences()
	preferences.IncludeCompletionsForModuleExports = core.TSFalse
	preferences.IncludeCompletionsForImportStatements = core.TSTrue
	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "I",
					InsertText: new(`import { I } from "./a";`),
					Data: &lsproto.CompletionItemData{
						AutoImport: &lsproto.AutoImportFix{ModuleSpecifier: "./a"},
					},
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: `import { I } from "./a";`,
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
		UserPreferences: &preferences,
	})
}
