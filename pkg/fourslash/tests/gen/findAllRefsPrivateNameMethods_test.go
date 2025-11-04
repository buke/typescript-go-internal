package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsPrivateNameMethods(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C {
    /*1*/#foo(){ }
    constructor() {
        this./*2*/#foo();
    }
}
class D extends C {
    constructor() {
        super()
        this.#foo = 20;
    }
}
class E {
    /*3*/#foo(){ }
    constructor() {
        this./*4*/#foo();
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3", "4")
}
