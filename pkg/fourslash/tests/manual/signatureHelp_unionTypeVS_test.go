package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestSignatureHelp_unionTypeVS(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare const a: (fn?: ((x: string) => string) | ((y: number) => number)) => void;
declare const b: (x: string | number) => void;

interface Callback {
    (x: string): string;
    (x: number): number;
    (x: string | number): string | number;
}
declare function c(callback: Callback): void;
a((/*1*/) => {
    return undefined;
});

b(/*2*/);

c((/*3*/) => {});`
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}
