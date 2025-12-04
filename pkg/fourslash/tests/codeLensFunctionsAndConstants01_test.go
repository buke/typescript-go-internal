package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCodeLensFunctionsAndConstants01(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `
// @module: preserve

// @filename: ./exports.ts

let callCount = 0;
export function foo(n: number): void {
  callCount++;
  if (n > 0) {
	foo(n - 1);
  }
  else {
    console.log("function was called " + callCount + " times");
  }
}

foo(5);

export const bar = 123;

// @filename: ./importer.ts
import { foo, bar } from "./exports";

foo(5);
console.log(bar);
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineCodeLens(t, &lsutil.UserPreferences{
		CodeLens: lsutil.CodeLensUserPreferences{
			ReferencesCodeLensEnabled:            true,
			ReferencesCodeLensShowOnAllFunctions: true,

			ImplementationsCodeLensEnabled:                true,
			ImplementationsCodeLensShowOnInterfaceMethods: true,
			ImplementationsCodeLensShowOnAllClassMethods:  true,
		},
	})
}
