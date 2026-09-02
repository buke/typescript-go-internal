package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestQuickinfoVerbosityNamespaceDefaultExport(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
declare namespace ns/*1*/ {
    interface Shape {
        sides: number;
    }
    const circle: Shape;
    export default circle;
    export { Shape };
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHoverWithVerbosity(t, map[string][]int{"1": {0, 1}})
}
