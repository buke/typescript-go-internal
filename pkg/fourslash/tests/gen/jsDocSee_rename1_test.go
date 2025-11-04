package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsDocSee_rename1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|interface [|{| "contextRangeIndex": 0 |}A|] {}|]
/**
 * @see {[|A|]}
 */
declare const a: [|A|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, ToAny(f.Ranges()[1:])...)
}
