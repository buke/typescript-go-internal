package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestCallHierarchyIncomingCallsObjectLiteralMethodInIdentifierComputedProperty(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const key = "x";
const obj = {
  [key]: {
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
