package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameTemplateLiteralsDefinePropertyJs(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: a.js
let obj = {};

Object.defineProperty(obj, ` + "`" + `[|prop|]` + "`" + `, { value: 0 });

obj = {
    [|[` + "`" + `[|{| "contextRangeIndex": 1 |}prop|]` + "`" + `]: 1|]
};

obj.[|prop|];
obj['[|prop|]'];
obj["[|prop|]"];
obj[` + "`" + `[|prop|]` + "`" + `];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "prop")
}
