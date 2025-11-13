package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixNewImportNodeModules1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|f1/*0*/();|]
// @Filename: ../package.json
{ "dependencies": { "fake-module": "latest" } }
// @Filename: ../node_modules/fake-module/nested.ts
export var v1 = 5;
export function f1();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyImportFixAtPosition(t, []string{
		`import { f1 } from "fake-module/nested";

f1();`,
	}, nil /*preferences*/)
}
