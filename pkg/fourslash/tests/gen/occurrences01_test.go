package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestOccurrences01(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `foo: [|switch|] (10) {
    [|case|] 1:
    [|case|] 2:
    [|case|] 3:
        [|break|];
        [|break|] foo;
        continue;
        continue foo;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
