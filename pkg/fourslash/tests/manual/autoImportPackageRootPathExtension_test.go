package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestAutoImportPackageRootPathExtension(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: /node_modules/pkg/package.json
{
    "name": "pkg",
    "version": "1.0.0",
    "main": "lib"
 }
// @Filename: /node_modules/pkg/lib/index.d.mts
export declare function foo(): any;
// @Filename: /package.json
{
    "dependencies": {
       "pkg": "*"
    }
 }
// @Filename: /index.ts
foo/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{AutoImportEntrypointDirectorySearch: core.TSTrue})
	f.VerifyImportFixModuleSpecifiers(t, "", []string{"pkg/lib/index.mjs"}, nil /*preferences*/)
}
