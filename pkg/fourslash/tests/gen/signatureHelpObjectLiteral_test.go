package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpObjectLiteral(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var objectLiteral = { n: 5, s: "", f: (a: number, b: string) => "" };
objectLiteral.f(/*objectLiteral1*/4, /*objectLiteral2*/"");`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "objectLiteral1")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "f(a: number, b: string): string", ParameterCount: 2, ParameterName: "a", ParameterSpan: "a: number"})
	f.GoToMarker(t, "objectLiteral2")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "f(a: number, b: string): string", ParameterName: "b", ParameterSpan: "b: string"})
}
