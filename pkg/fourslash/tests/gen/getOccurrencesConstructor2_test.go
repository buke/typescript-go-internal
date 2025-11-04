package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesConstructor2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C {
    constructor();
    constructor(x: number);
    constructor(y: string, x: number);
    constructor(a?: any, ...r: any[]) {
        if (a === undefined && r.length === 0) {
            return;
        }

        return;
    }
}

class D {
    [|con/**/structor|](public x: number, public y: number) {
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
