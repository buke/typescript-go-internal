package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoLink11(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**
 * {@link https://vscode.dev}
 * [link text]{https://vscode.dev}
 * {@link https://vscode.dev|link text}
 * {@link https://vscode.dev link text}
 */
function f() {}

/**/f();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
