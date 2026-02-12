package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionColonToken(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `
// @filename: /a.ts
:/*a*/

// @filename: /b.ts
function b(class: /*b*/) {}

// @filename: /c.ts
function c(enum: /*c*/) {}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	for _, marker := range f.Ranges() {
		f.VerifyCompletions(t, marker, &fourslash.CompletionsExpectedList{
			IsIncomplete: false,
			ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
				CommitCharacters: &DefaultCommitCharacters,
				EditRange:        Ignored,
			},
			Items: &fourslash.CompletionsExpectedItems{
				Includes: CompletionGlobals,
			},
		})
	}
}
