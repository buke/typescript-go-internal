package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestInlayHintsTupleTypeCrash(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function iterateTuples(tuples: [string][]): void {
  tuples.forEach((l) => {})
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayFunctionParameterTypeHints: core.TSTrue,
		},
	})
}
