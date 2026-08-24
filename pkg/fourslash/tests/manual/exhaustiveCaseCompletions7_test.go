package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestExhaustiveCaseCompletions7(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @newline: LF
export function foo(position: -1 | 0 | 1) {
    switch (position) {
        /**/
    }
}`
	f, done := fourslash.NewFourslash(t, fourslash.GetDefaultCapabilitiesWithOptions(&fourslash.ClientCapabilitiesOptions{
		CompletionItem: &lsproto.ClientCompletionItemOptions{
			SnippetSupport: new(true),
		},
	}), content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:            "case -1: ...",
					InsertText:       new("case -1:$1\ncase 0:$2\ncase 1:$3"),
					SortText:         new(string(ls.SortTextGlobalsOrKeywords)),
					InsertTextFormat: new(lsproto.InsertTextFormatSnippet),
				},
			},
		},
	})
}
