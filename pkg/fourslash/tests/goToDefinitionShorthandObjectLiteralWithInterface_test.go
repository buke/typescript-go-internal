package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionShorthandObjectLiteralWithInterface(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Something {
    [|foo|]: string;
}

function makeSomething([|foo|]: string): Something {
    return { [|f/*1*/oo|] };
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, true, "1")
}
