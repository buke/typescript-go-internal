package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickinfoVerbosityNamespaceErrorClassHeritage1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
namespace NS/*1*/ {
    export class Derived extends NonExistentClass {
        derivedField: number;
    }
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHoverWithVerbosity(t, map[string][]int{"1": {0, 1}})
}
