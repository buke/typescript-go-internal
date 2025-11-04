package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoOnExpandoLikePropertyWithSetterDeclarationJs1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
// @checkJs: true
// @filename: index.js
const x = {};

Object.defineProperty(x, "foo", {
  /** @param {number} v */
  set(v) {},
});

x.foo/**/ = 1;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "(property) x.foo: number", "")
}
