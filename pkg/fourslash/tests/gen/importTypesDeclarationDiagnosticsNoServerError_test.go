package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportTypesDeclarationDiagnosticsNoServerError(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @declaration: true
// @Filename: node_modules/foo/index.d.ts
export function f(): I;
export interface I {
  x: number;
}
// @Filename: a.ts
import { f } from "foo";
export const x = f();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFileNumber(t, 1)
	f.VerifyNonSuggestionDiagnostics(t, nil)
}
