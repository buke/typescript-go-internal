package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsJSDocNamespacedTypedef(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
// @allowJs: true
// @checkJs: true
// @Filename: /index.js
// Namespaced typedef
/** @typedef {string} [|NS|].[|T|] */

// Namespaced typedef aliased to qualified namespaced typedef.
/** @typedef {NS.T} NS.[|U|] */

// Namespaced typedef aliased to implicitly-resolved typedef.
/** @typedef {U} NS.[|V|] */
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t)
}
