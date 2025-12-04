package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestProcessInvalidSyntax1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: decl.js
var obj = {};
// @Filename: unicode1.js
obj.𝒜 ;
// @Filename: unicode2.js
obj.¬ ;
// @Filename: unicode3.js
obj¬
// @Filename: forof.js
for (obj/**/.prop of arr) {

}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, nil /*preferences*/, "")
}
