package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameNumericalIndexSingleQuoted(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const foo = { [|0|]: true };
foo[[|0|]];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, &lsutil.UserPreferences{QuotePreference: lsutil.QuotePreference("single")}, "0")
}
