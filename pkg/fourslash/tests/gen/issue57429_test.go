package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestIssue57429(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
function Builder<I>(def: I) {
  return def;
}

interface IThing {
  doThing: (args: { value: object }) => string
  doAnotherThing: () => void
}

Builder<IThing>({
  doThing(args: { value: object }) {
    const { v/*1*/alue } = this.[|args|]
    return ` + "`" + `${value}` + "`" + `
  },
  doAnotherThing() { },
})`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "const value: any", "")
	f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{
		{
			Message: "Property 'args' does not exist on type 'IThing'.",
			Code:    &lsproto.IntegerOrString{Integer: PtrTo[int32](2339)},
		},
	})
}
