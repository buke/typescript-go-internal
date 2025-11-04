package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestContextualTypingOfGenericCallSignatures1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var f24: {
   <T, U>(x: T): U
};
// x should not be contextually typed 
var f24 = (/**/x) => { return 1 };`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "(parameter) x: any", "")
}
