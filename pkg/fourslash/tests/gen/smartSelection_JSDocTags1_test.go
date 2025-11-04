package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSmartSelection_JSDocTags1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**
 * @returns {Array<{ value: /**/string }>}
 */
function foo() { return [] }`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineSelectionRanges(t)
}
