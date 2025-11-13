package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportFixWithMultipleModuleExportAssignment(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: esnext
// @allowJs: true
// @checkJs: true
// @Filename: /a.js
function f() {}
module.exports = f;
module.exports = 42;
// @Filename: /b.js
export const foo = 0;
// @Filename: /c.js
foo`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/c.js")
	f.VerifyImportFixAtPosition(t, []string{
		`const { foo } = require("./b");

foo`,
	}, nil /*preferences*/)
}
