package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionListsStringLiteralTypeAsIndexedAccessTypeObject(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `let firstCase: "a/*case_1*/"["foo"]
let secondCase: "b/*case_2*/"["bar"]
let thirdCase: "c/*case_3*/"["baz"]
let fourthCase: "en/*case_4*/"["qux"]
interface Foo {
 bar: string;
 qux: string;
}
let fifthCase: Foo["b/*case_5*/"]
let sixthCase: Foo["qu/*case_6*/"]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"case_1", "case_2", "case_3", "case_4"}, nil)
	f.VerifyCompletions(t, "case_5", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "bar",
				},
			},
		},
	})
	f.VerifyCompletions(t, "case_6", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "qux",
				},
			},
		},
	})
}
