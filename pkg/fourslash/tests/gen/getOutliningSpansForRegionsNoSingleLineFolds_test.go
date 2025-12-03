package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOutliningSpansForRegionsNoSingleLineFolds(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|//#region
function foo()[| {

}|]
[|//these
//should|]
//#endregion not you|]
[|// be
// together|]

[|//#region bla bla bla

function bar()[| { }|]

//#endregion|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.MarkTestAsStradaServer()
	f.VerifyOutliningSpans(t)
}
