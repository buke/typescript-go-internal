package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestIncrementalResolveConstructorDeclaration(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class c1 {
    private b: number;
    constructor(a: string) {
        this.b = a;
    }
}
var val = new c1("hello");
/*1*/val;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "var val: c1", "")
	f.VerifyNumberOfErrorsInCurrentFile(t, 1)
}
