package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/ls"
	"github.com/buke/typescript-go-internal/v7/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestCompletionsSelfDeclaring1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Test {
  keyPath?: string;
  autoIncrement?: boolean;
}

function test<T extends Record<string, Test>>(opt: T) { }

test({
  a: {
    keyPath: '',
    [|a|]/**/
  }
})`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "autoIncrement?",
					FilterText: new("autoIncrement"),
					SortText:   new(string(ls.SortTextOptionalMember)),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						InsertReplaceEdit: &lsproto.InsertReplaceEdit{
							NewText: "autoIncrement",
							Insert:  f.Ranges()[0].LSRange,
							Replace: f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
}
