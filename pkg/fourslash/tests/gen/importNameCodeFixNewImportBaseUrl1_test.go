package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixNewImportBaseUrl1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /tsconfig.json
{
    "compilerOptions": {
        "baseUrl": "./a"
    }
}
// @Filename: /a/b/x.ts
export function f1() { };
// @Filename: /a/b/y.ts
[|f1/*0*/();|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/a/b/y.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import { f1 } from "./x";

f1();`,
	}, nil /*preferences*/)
	f.VerifyImportFixAtPosition(t, []string{
		`import { f1 } from "b/x";

f1();`,
	}, nil /*preferences*/)
}
