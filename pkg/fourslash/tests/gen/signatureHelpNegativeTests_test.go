package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpNegativeTests(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//inside a comment foo(/*insideComment*/
cl/*invalidContext*/ass InvalidSignatureHelpLocation { }
InvalidSignatureHelpLocation(/*validContext*/);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyNoSignatureHelpForMarkers(t, "insideComment", "invalidContext", "validContext")
}
