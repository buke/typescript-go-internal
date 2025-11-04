package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionVariableAssignment1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @checkJs: true
// @filename: foo.js
const Foo = module./*def*/exports = function () {}
Foo.prototype.bar = function() {}
new [|Foo/*ref*/|]();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "foo.js")
	f.VerifyBaselineGoToDefinition(t, true, "ref")
}
