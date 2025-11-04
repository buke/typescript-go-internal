package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsForComputedProperties(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I {
    ["/*0*/prop1"]: () => void;
}

class C implements I {
    ["/*1*/prop1"]: any;
}

var x: I = {
    ["/*2*/prop1"]: function () { },
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "0", "1", "2")
}
