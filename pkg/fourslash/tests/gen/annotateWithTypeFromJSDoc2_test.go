package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestAnnotateWithTypeFromJSDoc2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: test123.ts
/** @type {number} */
var [|x|]: string;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifySuggestionDiagnostics(t, nil)
}
