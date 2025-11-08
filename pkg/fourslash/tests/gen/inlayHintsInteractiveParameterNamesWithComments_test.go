package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestInlayHintsInteractiveParameterNamesWithComments(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const fn = (x: any) => { }
fn(/* nobody knows exactly what this param is */ 42);
function foo (aParameter: number, bParameter: number, cParameter: number) { }
foo(
    /** aParameter */
    1,
    // bParameter
    2,
    /* cParameter */
    3
)
foo(
    /** multiple comments */
    /** aParameter */
    1,
    /** bParameter */
    /** multiple comments */
    2,
    // cParameter
    /** multiple comments */
    3
)
foo(
    /** wrong name */
    1,
    2,
    /** multiple */
    /** wrong */
    /** name */
    3
)`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{IncludeInlayParameterNameHints: lsutil.IncludeInlayParameterNameHintsLiterals})
}
