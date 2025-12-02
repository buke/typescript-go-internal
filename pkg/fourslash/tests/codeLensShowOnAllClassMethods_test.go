package fourslash_test

import (
	"fmt"
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCodeLensReferencesShowOnAllClassMethods(t *testing.T) {
	t.Parallel()
	containingTestName := t.Name()
	for _, value := range []bool{true, false} {
		t.Run(fmt.Sprintf("%s=%v", containingTestName, value), func(t *testing.T) {
			t.Parallel()
			defer testutil.RecoverAndFail(t, "Panic on fourslash test")

			const content = `
export abstract class ABC {
  abstract methodA(): void;
  methodB(): void {}
  #methodC(): void {}
  protected methodD(): void {}
  private methodE(): void {}
  protected abstract methodG(): void;
  public methodH(): void {}

  static methodStaticA(): void {}
  protected static methodStaticB(): void {}
  private static methodStaticC(): void {}
  static #methodStaticD(): void {}
}
`
			f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
			f.VerifyBaselineCodeLens(t, &lsutil.UserPreferences{
				ImplementationsCodeLensEnabled:               true,
				ImplementationsCodeLensShowOnAllClassMethods: value,
			})
		})
	}
}
