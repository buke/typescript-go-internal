package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCommentsUnion(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var a: Array<string> | Array<number>;
a./*1*/length`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "(property) Array<T>.length: number", "Gets or sets the length of the array. This is a number one higher than the highest index in the array.")
}
