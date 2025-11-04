package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsGenericTypeWithMultipleBases1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `export interface iBaseScope {
    watch: () => void;
}
export interface iMover {
    moveUp: () => void;
}
export interface iScope<TModel> extends iBaseScope, iMover {
    family: TModel;
}
var x: iScope<number>;
x./**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "family",
					Detail: PtrTo("(property) iScope<number>.family: number"),
				},
				&lsproto.CompletionItem{
					Label:  "moveUp",
					Detail: PtrTo("(property) iMover.moveUp: () => void"),
				},
				&lsproto.CompletionItem{
					Label:  "watch",
					Detail: PtrTo("(property) iBaseScope.watch: () => void"),
				},
			},
		},
	})
}
