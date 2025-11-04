package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsSelfDeclaring2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f1<T>(x: T) {}
f1({ [|abc|]/*1*/ });`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange: &fourslash.EditRange{
				Insert:  f.Ranges()[0],
				Replace: f.Ranges()[0],
			},
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: CompletionGlobalsPlus([]fourslash.CompletionsExpectedItem{
				"f1",
			}, false /*noLib*/),
		},
	})
}
