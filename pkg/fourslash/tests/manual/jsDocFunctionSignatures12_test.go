package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsDocFunctionSignatures12(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: jsDocFunctionSignatures.js
/**
 * @param {{
 *   stringProp: string,
 *   numProp: number,
 *   boolProp: boolean,
 *   anyProp: any,
 *   anotherAnyProp: any,
 *   functionProp: (arg0: string, arg1: any) => any
 * }} o
 */
function f1(o) {
    o/**/;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.VerifyQuickInfoIs(t, "(parameter) o: { stringProp: string; numProp: number; boolProp: boolean; anyProp: any; anotherAnyProp: any; functionProp: (arg0: string, arg1: any) => any; }", "")
}
