package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionForStringLiteral_quotePreference8(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: /a.ts
export const a = null;
// @filename: /b.ts
import { a } from './a';

const foo = { '"a name\'s all good but it\'s better with more"': null };
foo[|./**/|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/b.ts")
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "\"a name's all good but it's better with more\"",
					InsertText: PtrTo("['\"a name\\'s all good but it\\'s better with more\"']"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "\"a name's all good but it's better with more\"",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
}
