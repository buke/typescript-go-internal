package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsDocPropertyDescription3(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface LiteralExample {
    /** Something generic */
    [key: ` + "`" + `data-${string}` + "`" + `]: string;
     /** Something else */
    [key: ` + "`" + `prefix${number}` + "`" + `]: number;
}
function literalExample(e: LiteralExample) {
    console.log(e./*literal*/anything);
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "literal", "any", "")
}
