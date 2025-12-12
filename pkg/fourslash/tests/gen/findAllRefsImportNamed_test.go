package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsImportNamed(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: f.ts
export { foo as foo }
function /*start*/foo(a: number, b: number) { }
// @Filename: b.ts
import x = require("./f");
x.foo(1, 2);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyNoErrors(t)
	f.VerifyBaselineFindAllReferences(t, "start")
}
