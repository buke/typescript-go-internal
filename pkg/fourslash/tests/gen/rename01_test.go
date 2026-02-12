package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRename01(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: es5
///<reference path="./Bar.ts" />
[|function [|{| "contextRangeIndex": 0 |}Bar|]() {
    // This is a reference to [|Bar|] in a comment.
    "this is a reference to [|Bar|] in a string"
}|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.MarkTestAsStradaServer()
	f.VerifyBaselineRename(t, nil /*preferences*/, f.Ranges()[1])
}
