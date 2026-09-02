package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestInlayHintsPropertyDeclarationComputedName1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo() {
  const sym = Symbol();
  class C {
    [sym] = 123;
  }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayPropertyDeclarationTypeHints: core.TSTrue,
		},
	})
}
