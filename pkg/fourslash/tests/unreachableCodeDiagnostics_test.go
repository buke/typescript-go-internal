package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestUnreachableCodeDiagnostics(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowUnreachableCode: false
throw new Error();
	
(() => {})();
	`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineNonSuggestionDiagnostics(t)
}
