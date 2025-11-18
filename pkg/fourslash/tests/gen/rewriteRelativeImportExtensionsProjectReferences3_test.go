package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestRewriteRelativeImportExtensionsProjectReferences3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: src/tsconfig-base.json
{
    "compilerOptions": {
        "module": "nodenext",
        "composite": true,
        "rewriteRelativeImportExtensions": true,
    }
}
// @Filename: src/compiler/tsconfig.json
{
    "extends": "../tsconfig-base.json",
    "compilerOptions": {
        "rootDir": ".",
        "outDir": "../../dist/compiler",
}
// @Filename: src/compiler/parser.ts
export {};
// @Filename: src/services/tsconfig.json
{
    "extends": "../tsconfig-base.json",
    "compilerOptions": {
        "rootDir": ".",
        "outDir": "../../dist/services",
    },
    "references": [
        { "path": "../compiler" }
    ]
}
// @Filename: src/services/services.ts
import {} from "../compiler/parser.ts";`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/src/services/services.ts")
	f.VerifyBaselineNonSuggestionDiagnostics(t)
}
