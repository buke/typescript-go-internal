package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNavigationBarFunctionPrototypeNested(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: foo.js
function A() {}
A.B = function () {  } 
A.B.prototype.d = function () {  }  
Object.defineProperty(A.B.prototype, "x", {
    get() {}
})
A.prototype.D = function () {  } 
A.prototype.D.prototype.d = function () {  } `
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
