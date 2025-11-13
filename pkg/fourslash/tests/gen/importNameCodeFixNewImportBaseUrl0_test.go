package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixNewImportBaseUrl0(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|f1/*0*/();|]
// @Filename: tsconfig.json
{
    "compilerOptions": {
        "baseUrl": "./a"
    }
}
// @Filename: a/b.ts
export function f1() { };`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyImportFixAtPosition(t, []string{
		`import { f1 } from "b";

f1();`,
	}, nil /*preferences*/)
}
