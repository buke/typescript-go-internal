package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoOnPropDeclaredUsingIndexSignatureOnInterfaceWithBase(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface P {}
interface B extends P {
  [k: string]: number;
}
declare const b: B;
b.t/*1*/est = 10;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "(index) B[string]: number", "")
}
