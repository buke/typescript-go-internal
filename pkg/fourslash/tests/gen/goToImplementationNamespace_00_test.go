package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToImplementationNamespace_00(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `namespace /*implementation0*/Foo {
    export function hello() {}
}

module /*implementation1*/Bar {
    export function sure() {}
}

let x = Fo/*reference0*/o;
let y = Ba/*reference1*/r;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToImplementation(t, "reference0", "reference1")
}
