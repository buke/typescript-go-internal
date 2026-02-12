package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetDeclarationDiagnostics(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: false
// @declaration: true
// @outFile: true
// @Filename: inputFile1.ts
namespace m {
   export function foo() {
       class C implements I { private a; }
       interface I { }
       return C;
   }
} /*1*/
// @Filename: input2.ts
var x = "hello world"; /*2*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.VerifyNumberOfErrorsInCurrentFile(t, 1)
	f.GoToMarker(t, "2")
	f.VerifyNumberOfErrorsInCurrentFile(t, 0)
}
