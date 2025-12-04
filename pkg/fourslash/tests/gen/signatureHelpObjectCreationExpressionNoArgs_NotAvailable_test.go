package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpObjectCreationExpressionNoArgs_NotAvailable(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class sampleCls { constructor(str: string, num: number) { } }
var x = new sampleCls/**/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyNoSignatureHelpForMarkers(t, "")
}
