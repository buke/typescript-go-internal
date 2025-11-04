package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllReferencesJsOverloadedFunctionParameter(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @checkJs: true
// @Filename: foo.js
/**
 * @overload
 * @param {number} x
 * @returns {number}
 *
 * @overload
 * @param {string} x
 * @returns {string} 
 *
 * @param {unknown} x
 * @returns {unknown} 
 */
function foo(x/*1*/) {
  return x;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1")
}
