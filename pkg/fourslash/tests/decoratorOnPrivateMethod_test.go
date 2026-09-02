package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestDecoratorCompletionOnPrivateMethod(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
// @experimentalDecorators: true
declare function dec(target: any, key: string): void;
class C {
    @dec/**/
    #method() {}
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	// Verify completions don't panic on decorator applied to private method
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"dec",
			},
		},
	})
}
