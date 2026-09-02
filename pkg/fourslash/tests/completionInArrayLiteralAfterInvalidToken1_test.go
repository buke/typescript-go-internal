package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestCompletionInArrayLiteralAfterInvalidToken1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `const pairs: Record<string, [string, string]> = {
  a: ["x",:/*m1*/]
};
`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCompletions(t, "m1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{},
	})
}
