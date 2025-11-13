package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_importType(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @checkJs: true
// @Filename: /a.js
export {};
/** @typedef {number} T */
// @Filename: /b.js
/** @type {T} */
const x = 0;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/b.js")
	f.VerifyImportFixAtPosition(t, []string{
		`/** @type {import("./a").T} */
const x = 0;`,
	}, nil /*preferences*/)
}
