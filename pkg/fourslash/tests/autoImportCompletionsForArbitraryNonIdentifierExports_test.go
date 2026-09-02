package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	util "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestAutoImportCompletionsForArbitraryNonIdentifierExports(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
// @module: esnext
// @Filename: /a.ts
const foo = 0;
export { foo as "foo-bar" };
export const fooBar = 1;

// @Filename: /b.ts
foo/**/
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		Items: &fourslash.CompletionsExpectedItems{
			Excludes: []string{"foo-bar"},
			Includes: []fourslash.CompletionsExpectedItem{"fooBar"},
		},
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &util.DefaultCommitCharacters,
			EditRange:        util.Ignored,
		},
	})
}
