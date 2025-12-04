package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionJsModuleNameAtImportName(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: /foo.js
 /*moduleDef*/function notExported() { }
 class Blah {
    abc = 123;
 }
 module.exports.Blah = Blah;
// @Filename: /bar.js
const [|/*importDef*/BlahModule|] = require("./foo.js");
new [|/*importUsage*/BlahModule|].Blah()
// @Filename: /barTs.ts
import [|/*importDefTs*/BlahModule|] = require("./foo.js");
new [|/*importUsageTs*/BlahModule|].Blah()`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, true, "importDef", "importUsage", "importDefTs", "importUsageTs")
}
