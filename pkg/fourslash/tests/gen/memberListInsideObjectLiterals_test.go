package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestMemberListInsideObjectLiterals(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module ObjectLiterals {
    interface MyPoint {
        x1: number;
        y1: number;
    }

    var p1: MyPoint = {
        /*1*/
    };

    var p2: MyPoint = {
        x1: 5,
        /*2*/
    };

    var p3: MyPoint = {
        x1/*3*/:
    };

    var p4: MyPoint = {
        /*4*/y1
    };
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"1", "3", "4"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "x1",
					Detail: PtrTo("(property) MyPoint.x1: number"),
				},
				&lsproto.CompletionItem{
					Label:  "y1",
					Detail: PtrTo("(property) MyPoint.y1: number"),
				},
			},
		},
	})
	f.VerifyCompletions(t, []string{"2"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "y1",
					Detail: PtrTo("(property) MyPoint.y1: number"),
				},
			},
		},
	})
}
