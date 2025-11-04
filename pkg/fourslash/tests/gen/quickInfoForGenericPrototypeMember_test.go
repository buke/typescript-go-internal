package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoForGenericPrototypeMember(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C<T> {
   foo(x: T) { }
}
var x = new /*1*/C<any>();
var y = C.proto/*2*/type;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "constructor C<any>(): C<any>", "")
	f.VerifyQuickInfoAt(t, "2", "(property) C<T>.prototype: C<any>", "")
}
