package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRenameNamedImportUseAliasesForRenames(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
import { /*import*/MyTypeA } from "./b";
const type: MyTypeA = { foo: "bar" };
// @Filename: /b.ts
export interface MyTypeA {
    foo: string;
}`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, &lsutil.UserPreferences{UseAliasesForRename: core.TSFalse}, "import")
	f.VerifyBaselineRename(t, &lsutil.UserPreferences{UseAliasesForRename: core.TSTrue}, "import")
}

func TestRenameNamedImportDefaultInNodeModules(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /index.ts
import { /*fooImport*/[|Foo|] } from "foo";
declare const f: Foo;
// @Filename: /tsconfig.json
{}
// @Filename: /node_modules/foo/package.json
{ "types": "index.d.ts" }
// @Filename: /node_modules/foo/index.d.ts
export interface Foo {
    bar: string;
}`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, nil /*preferences*/, "fooImport")
	f.VerifyBaselineRename(t, &lsutil.UserPreferences{UseAliasesForRename: core.TSTrue}, "fooImport")
	f.GoToMarker(t, "fooImport")
	f.VerifyRenameFailed(t, &lsutil.UserPreferences{UseAliasesForRename: core.TSFalse})
}
