package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameJsPropertyAssignment4(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: /a.js
function f() {
   var /*1*/foo = this;
   /*2*/foo.x = 1;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/a.js")
	f.VerifyBaselineRename(t, nil /*preferences*/, "1", "2")
}
