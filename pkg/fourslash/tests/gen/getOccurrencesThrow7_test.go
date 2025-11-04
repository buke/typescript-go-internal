package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesThrow7(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `try {
    [|throw|] 10;

    try {
        throw 10;
    }
    catch (x) {
        [|throw|] 10;
    }
    finally {
        [|throw|] 10;
    }
}
finally {
    [|throw|] 10;
}

[|throw|] 10;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
