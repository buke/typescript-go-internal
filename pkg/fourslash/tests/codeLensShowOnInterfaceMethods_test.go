package fourslash_test

import (
	"fmt"
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCodeLensReferencesShowOnInterfaceMethods(t *testing.T) {
	t.Parallel()
	containingTestName := t.Name()
	for _, value := range []core.Tristate{core.TSTrue, core.TSFalse} {
		t.Run(fmt.Sprintf("%s=%v", containingTestName, value.IsTrue()), func(t *testing.T) {
			t.Parallel()
			defer testutil.RecoverAndFail(t, "Panic on fourslash test")

			const content = `
export interface I {
  methodA(): void;
}
export interface I {
  methodB(): void;
}

interface J extends I {
  methodB(): void;
  methodC(): void;
}

class C implements J {
  methodA(): void {}
  methodB(): void {}
  methodC(): void {}
}

class AbstractC implements J {
  abstract methodA(): void;
  methodB(): void {}
  abstract methodC(): void;
}
`
			f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
			defer done()
			f.VerifyBaselineCodeLens(t, &lsutil.UserPreferences{
				CodeLens: lsutil.CodeLensUserPreferences{
					ImplementationsCodeLensEnabled:                core.TSTrue,
					ImplementationsCodeLensShowOnInterfaceMethods: value,
				},
			})
		})
	}
}
