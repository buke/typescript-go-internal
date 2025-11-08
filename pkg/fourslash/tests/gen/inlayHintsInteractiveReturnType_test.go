package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestInlayHintsInteractiveReturnType(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo1 () {
    return 1
}
function foo2 (): number {
    return 1
}
class C {
    foo() {
        return 1
    }
    bar() {
        return this
    }
}
const a = () => 1
const b = function () { return 1 }
const c = (b) => 1
const d = b => 1`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{IncludeInlayFunctionLikeReturnTypeHints: true})
}
