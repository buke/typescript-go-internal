package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameObjectSpread(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface A1 { [|[|{| "contextRangeIndex": 0 |}a|]: number|] };
interface A2 { [|[|{| "contextRangeIndex": 2 |}a|]?: number|] };
let a1: A1;
let a2: A2;
let a12 = { ...a1, ...a2 };
a12.[|a|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1], f.Ranges()[3], f.Ranges()[4])
}
