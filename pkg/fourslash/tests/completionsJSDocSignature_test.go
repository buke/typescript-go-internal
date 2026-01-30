package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsJSDocSignature(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `// @noLib: true
// @checkJs: true
// @allowJs: true
// @filename: index.js
/**
 * @type {{
 *   (input: string):/*1*/ X|Y/*2*/
 * }}
 */
let x;`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{".", ",", ";"},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{".", ",", ";"},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{},
	})
}
