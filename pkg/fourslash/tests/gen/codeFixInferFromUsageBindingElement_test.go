package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCodeFixInferFromUsageBindingElement(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f([car, cdr]) {
    return car + cdr + 1
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifySuggestionDiagnostics(t, nil)
}
