package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsObjectBindingElementPropertyName10(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Recursive {
    /*1*/next?: Recursive;
    value: any;
}

function f (/*2*/{ /*3*/next: { /*4*/next: x} }: Recursive) {
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3", "4")
}
