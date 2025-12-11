package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixJsEnding(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @Filename: /node_modules/lit/package.json
{ "name": "lit", "version": "1.0.0" }
// @Filename: /node_modules/lit/index.d.ts
import "./decorators";
// @Filename: /node_modules/lit/decorators.d.ts
export declare function customElement(name: string): any;
// @Filename: /a.ts
customElement/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyImportFixModuleSpecifiers(t, "", []string{"lit/decorators.js"}, &lsutil.UserPreferences{ImportModuleSpecifierEnding: "js"})
}
