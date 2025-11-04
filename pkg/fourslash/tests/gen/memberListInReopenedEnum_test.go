package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestMemberListInReopenedEnum(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module M {
    enum E {
        A, B
    }
    enum E {
        C = 0, D
    }
    var x = E./*1*/
}`
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
					Label:  "A",
					Detail: PtrTo("(enum member) E.A = 0"),
				},
				&lsproto.CompletionItem{
					Label:  "B",
					Detail: PtrTo("(enum member) E.B = 1"),
				},
				&lsproto.CompletionItem{
					Label:  "C",
					Detail: PtrTo("(enum member) E.C = 0"),
				},
				&lsproto.CompletionItem{
					Label:  "D",
					Detail: PtrTo("(enum member) E.D = 1"),
				},
			},
		},
	})
}
