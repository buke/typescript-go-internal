package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFormattingInComment(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class A {
foo(              ); // /*1*/
}
function foo() {       var x;       } // /*2*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.Insert(t, ";")
	f.VerifyCurrentLineContentIs(t, "foo(              ); // ;")
	f.GoToMarker(t, "2")
	f.Insert(t, "}")
	f.VerifyCurrentLineContentIs(t, "function foo() {       var x;       } // }")
}
