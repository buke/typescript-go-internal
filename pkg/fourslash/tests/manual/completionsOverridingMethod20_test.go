package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsOverridingMethod20(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
// @newline: LF
abstract class AFoo {
    abstract bar(): Promise<void>;
}
class Foo extends AFoo {
    async [|b/*a*/|]
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:               "bar",
					TextEdit:            InsertReplaceTextEdit("async bar(): Promise<void> {\n}", f.Ranges()[0].LSRange),
					FilterText:          new("bar"),
					SortText:            new(string(ls.SortTextLocationPriority)),
					AdditionalTextEdits: fourslash.AnyTextEdits,
					Data: &lsproto.CompletionItemData{
						Source: "ClassMemberSnippet/",
					},
				},
			},
		},
		UserPreferences: &lsutil.UserPreferences{IncludeCompletionsWithClassMemberSnippets: core.TSTrue},
	})
	f.VerifyApplyCodeActionFromCompletion(t, new("a"), &fourslash.ApplyCodeActionFromCompletionOptions{
		UserPreferences: &lsutil.UserPreferences{IncludeCompletionsWithClassMemberSnippets: core.TSTrue},
		Name:            "bar",
		Source:          "ClassMemberSnippet/",
		Description:     "Update modifiers of 'bar'",
		NewFileContent: new(`abstract class AFoo {
    abstract bar(): Promise<void>;
}
class Foo extends AFoo {
    b
}`),
	})
}
