package fourslash_test

import (
	"fmt"
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCodeLensReferencesShowOnAllFunctions(t *testing.T) {
	t.Parallel()
	containingTestName := t.Name()
	for _, value := range []bool{true, false} {
		t.Run(fmt.Sprintf("%s=%v", containingTestName, value), func(t *testing.T) {
			t.Parallel()
			defer testutil.RecoverAndFail(t, "Panic on fourslash test")

			const content = `
export function f1(): void {}

function f2(): void {}

export const f3 = () => {};

const f4 = () => {};

const f5 = function() {};
`
			f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
			f.VerifyBaselineCodeLens(t, &lsutil.UserPreferences{
				ReferencesCodeLensEnabled:            true,
				ReferencesCodeLensShowOnAllFunctions: value,
			})
		})
	}
}
