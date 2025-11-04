package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesReturn(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f(a: number) {
    if (a > 0) {
        [|ret/**/urn|] (function () {
            return;
            return;
            return;

            if (false) {
                return true;
            }
        })() || true;
    }

    var unusued = [1, 2, 3, 4].map(x => { return 4 })

    [|return|];
    [|return|] true;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
