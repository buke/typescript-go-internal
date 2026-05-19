package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestTrailingCommaSignatureHelpVS(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function str(n: number): string;
/**
 * Stringifies a number with radix
 * @param radix The radix
 */
function str(n: number, radix: number): string;
function str(n: number, radix?: number): string { return ""; }

str(1, /*a*/)

declare function f<T>(a: T): T;
f(2, /*b*/);`
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}
