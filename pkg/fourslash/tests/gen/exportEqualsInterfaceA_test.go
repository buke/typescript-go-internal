package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestExportEqualsInterfaceA(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: exportEqualsInterface_A.ts
interface A {
    p1: number;
}
export = A;
/**/
var i: I1;
var n: number = i.p1;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.Insert(t, "import I1 = require(\"exportEqualsInterface_A\");")
}
