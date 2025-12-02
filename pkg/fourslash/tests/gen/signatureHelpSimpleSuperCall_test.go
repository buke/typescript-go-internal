package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpSimpleSuperCall(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class SuperCallBase {
    constructor(b: boolean) {
    }
}
class SuperCall extends SuperCallBase {
    constructor() {
        super(/*superCall*/);
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "superCall")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "SuperCallBase(b: boolean): SuperCallBase", ParameterName: "b", ParameterSpan: "b: boolean"})
}
