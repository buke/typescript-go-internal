package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGenericObjectBaseType(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C<T> {
    constructor(){}
    foo(a: T) {
        return a.toString();
    }
}
var x = new C<string>();
var y: string = x.foo("hi");
/*1*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.VerifyNoErrors(t)
}
