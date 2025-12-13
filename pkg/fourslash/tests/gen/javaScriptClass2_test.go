package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJavaScriptClass2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowNonTsExtensions: true
// @Filename: Foo.js
class Foo {
   constructor() {
       [|this.[|{| "contextRangeIndex": 0 |}union|] = 'foo';|]
       [|this.[|{| "contextRangeIndex": 2 |}union|] = 100;|]
   }
   method() { return this.[|union|]; }
}
var x = new Foo();
x.[|union|];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "union")
}
