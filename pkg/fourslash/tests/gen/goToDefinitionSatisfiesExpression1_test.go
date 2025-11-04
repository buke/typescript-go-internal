package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionSatisfiesExpression1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const STRINGS = {
    [|/*definition*/title|]: 'A Title',
} satisfies Record<string,string>;

//somewhere in app
STRINGS.[|/*usage*/title|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "definition", "usage")
}
