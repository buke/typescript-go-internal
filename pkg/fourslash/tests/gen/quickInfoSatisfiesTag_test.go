package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoSatisfiesTag(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noEmit: true
// @allowJS: true
// @checkJs: true
// @filename: /a.js
/** @satisfies {number} comment */
const /*1*/a = 1;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
