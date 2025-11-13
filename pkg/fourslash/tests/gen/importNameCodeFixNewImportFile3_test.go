package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixNewImportFile3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|let t: XXX/*0*/.I;|]
// @Filename: ./module.ts
export module XXX {
   export interface I {
   }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyImportFixAtPosition(t, []string{
		`import { XXX } from "./module";

let t: XXX.I;`,
	}, nil /*preferences*/)
}
