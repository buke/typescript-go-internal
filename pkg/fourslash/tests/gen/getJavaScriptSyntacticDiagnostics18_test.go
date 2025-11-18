package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetJavaScriptSyntacticDiagnostics18(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: a.js
class C {
    x; // Regular property declaration allowed
    static y; // static allowed
    public z; // public not allowed
}
// @Filename: b.js
class C {
    x: number; // Types not allowed
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineNonSuggestionDiagnostics(t)
}
