package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionPropertyAssignment(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `export const /*FunctionResult*/Component = () => { return "OK"}
Component./*PropertyResult*/displayName = 'Component'

[|/*FunctionClick*/Component|]

Component.[|/*PropertyClick*/displayName|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "FunctionClick", "PropertyClick")
}
