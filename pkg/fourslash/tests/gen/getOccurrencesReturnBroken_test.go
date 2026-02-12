package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesReturnBroken(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `ret/*1*/urn;
retu/*2*/rn;
function f(a: number) {
    if (a > 0) {
        return (function () {
            () => [|return|];
            [|return|];
            [|return|];

            if (false) {
                [|return|] true;
            }
        })() || true;
    }

    var unusued = [1, 2, 3, 4].map(x => { return 4 })

    return;
    return true;
}

class A {
    ret/*3*/urn;
    r/*4*/eturn 8675309;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Markers())...)
}
