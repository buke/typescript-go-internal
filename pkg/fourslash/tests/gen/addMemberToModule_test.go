package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestAddMemberToModule(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module A {
    /*var*/
}
module /*check*/A {
    var p;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "check")
	f.VerifyQuickInfoExists(t)
	f.GoToMarker(t, "var")
	f.Insert(t, "var o;")
	f.GoToMarker(t, "check")
	f.VerifyQuickInfoExists(t)
}
