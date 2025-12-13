package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestAutoImportNodeNextJSRequire(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: node18
// @allowJs: true
// @checkJs: true
// @noEmit: true
// @Filename: /matrix.js
exports.variants = [];
// @Filename: /main.js
exports.dedupeLines = data => {
  variants/**/
}
// @Filename: /totally-irrelevant-no-way-this-changes-things-right.js
export default 0;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/main.js")
	f.VerifyImportFixAtPosition(t, []string{
		`const { variants } = require("./matrix")

exports.dedupeLines = data => {
  variants
}`,
	}, nil /*preferences*/)
}
