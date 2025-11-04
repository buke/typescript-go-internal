package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefs_importType_js4(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @allowJs: true
// @checkJs: true
// @Filename: /a.js
/**
 * @callback /**/A
 * @param {unknown} response
 */

module.exports = {};
// @Filename: /b.js
/** @typedef {import("./a").A} A */`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "")
}
