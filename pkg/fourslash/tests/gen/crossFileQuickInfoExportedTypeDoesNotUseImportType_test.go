package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCrossFileQuickInfoExportedTypeDoesNotUseImportType(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: b.ts
export interface B {}
export function foob(): {
    x: B,
    y: B
} {
    return null as any;
}
// @Filename: a.ts
import { foob } from "./b";
const thing/*1*/ = foob(/*2*/);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "const thing: {\n    x: B;\n    y: B;\n}", "")
	f.GoToMarker(t, "2")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "foob(): { x: B; y: B; }"})
}
