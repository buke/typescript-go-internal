package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameModuleExportsProperties3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: a.js
[|class [|{| "contextRangeIndex": 0 |}A|] {}|]
module.exports = { [|A|] }`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineRename(t, &lsutil.UserPreferences{UseAliasesForRename: core.TSTrue}, f.Ranges()[1], f.Ranges()[2])
}
