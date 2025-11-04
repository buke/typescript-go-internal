package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameAlias2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|module [|{| "contextRangeIndex": 0 |}SomeModule|] { export class SomeClass { } }|]
import M = [|SomeModule|];
import C = M.SomeClass;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "SomeModule")
}
