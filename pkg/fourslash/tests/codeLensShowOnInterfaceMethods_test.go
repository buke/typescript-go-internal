package fourslash_test

import (
	"fmt"
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCodeLensReferencesShowOnInterfaceMethods(t *testing.T) {
	t.Parallel()
	containingTestName := t.Name()
	for _, value := range []bool{true, false} {
		t.Run(fmt.Sprintf("%s=%v", containingTestName, value), func(t *testing.T) {
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
			f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
			f.VerifyBaselineCodeLens(t, &lsutil.UserPreferences{
				ImplementationsCodeLensEnabled:                true,
				ImplementationsCodeLensShowOnInterfaceMethods: value,
			})
		})
	}
}
