package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetJavaScriptQuickInfo1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowNonTsExtensions: true
// @Filename: Foo.js
/** @type {function(new:string,number)} */
var /**/v;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "var v: new (arg1: number) => string", "")
}
