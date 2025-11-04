package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsDocPropertyDescription5(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Multiple1Example {
    /** Something generic */
    [key: number | symbol | ` + "`" + `data-${string}` + "`" + ` | ` + "`" + `data-${number}` + "`" + `]: string;
}
function multiple1Example(e: Multiple1Example) {
    console.log(e./*multiple1*/anything);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "multiple1", "any", "")
}
