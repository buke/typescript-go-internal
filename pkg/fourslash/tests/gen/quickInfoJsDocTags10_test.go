package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoJsDocTags10(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noEmit: true
// @allowJs: true
// @Filename: quickInfoJsDocTags10.js
/**
 * @param {T1} a
 * @param {T2} a
 * @template T1,T2 Comment Text
 */
const /**/foo = (a, b) => {};`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
