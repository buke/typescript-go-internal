package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameStringLiteralTypes2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type Foo = "[|a|]" | "b";

class C {
    p: Foo = "[|a|]";
    m() {
        if (this.p === "[|a|]") {}
        if ("[|a|]" === this.p) {}

        if (this.p !== "[|a|]") {}
        if ("[|a|]" !== this.p) {}

        if (this.p == "[|a|]") {}
        if ("[|a|]" == this.p) {}

        if (this.p != "[|a|]") {}
        if ("[|a|]" != this.p) {}
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "a")
}
