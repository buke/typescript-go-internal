package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpOnOverloadsDifferentArity3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function f();
declare function f(s: string);
declare function f(s: string, b: boolean);
declare function f(n: number, b: boolean);

f(/**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "f(): any", ParameterCount: 0, OverloadsCount: 4})
	f.Insert(t, "x, ")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "f(s: string, b: boolean): any", ParameterCount: 2, ParameterName: "b", ParameterSpan: "b: boolean", OverloadsCount: 4})
	f.Insert(t, "x, ")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "f(s: string, b: boolean): any", ParameterCount: 2, OverloadsCount: 4})
}
