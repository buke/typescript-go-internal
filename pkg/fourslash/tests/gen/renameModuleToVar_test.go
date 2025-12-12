package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameModuleToVar(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface IMod {
    y: number;
}
declare module/**/ X: IMod;// {
//    export var y: numb;
var y: number;
module Y {
    var z = y + 5;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.Backspace(t, 6)
	f.Insert(t, "var")
	f.VerifyNoErrors(t)
}
