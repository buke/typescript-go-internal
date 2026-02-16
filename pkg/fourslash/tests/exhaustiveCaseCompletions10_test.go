package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestExhaustiveCaseCompletions10(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
declare const u: "$1" | "2";
switch (u) {
    case/*1*/
}`
	capabilities := fourslash.GetDefaultCapabilities()
	capabilities.TextDocument.Completion.CompletionItem.SnippetSupport = new(false)
	f, done := fourslash.NewFourslash(t, capabilities, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "case \"$1\": ...",
					InsertText: new("case \"$1\":\ncase \"2\":"),
					SortText:   new(string(ls.SortTextGlobalsOrKeywords)),
				},
			},
		},
	})
}
