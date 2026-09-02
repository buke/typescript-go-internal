package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestCallHierarchyIncomingCallsObjectLiteralMethodInExpressionComputedProperty(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const obj = {
  [1 + 2]: {
    method() {
      return ""./*split*/split(",");
    }
  }
};
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "split")
	f.VerifyBaselineCallHierarchy(t)
}
