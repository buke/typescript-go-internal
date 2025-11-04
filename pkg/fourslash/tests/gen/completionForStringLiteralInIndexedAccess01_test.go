package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionForStringLiteralInIndexedAccess01(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Foo {
    foo: string;
    bar: string;
}

let x: Foo["[|/*1*/|]"]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "bar",
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "bar",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
				&lsproto.CompletionItem{
					Label: "foo",
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "foo",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
}
