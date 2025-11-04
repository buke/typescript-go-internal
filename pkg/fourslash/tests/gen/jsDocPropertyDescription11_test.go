package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsDocPropertyDescription11(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type AliasExample = {
    /** Something generic */
    [p: string]: string;
    /** Something else */
    [key: ` + "`" + `any${string}` + "`" + `]: string;
}
function aliasExample(e: AliasExample) {
    console.log(e./*alias*/anything);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "alias", "(index) AliasExample[string | `any${string}`]: string", "Something generic\nSomething else")
}
