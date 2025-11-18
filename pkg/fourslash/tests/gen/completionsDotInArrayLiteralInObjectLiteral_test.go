package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsDotInArrayLiteralInObjectLiteral(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const o = { x: [[|.|][||]/**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{
		{
			Code:    &lsproto.IntegerOrString{Integer: PtrTo[int32](1109)},
			Message: "Expression expected.",
			Range:   f.Ranges()[0].LSRange,
		},
		{
			Code:    &lsproto.IntegerOrString{Integer: PtrTo[int32](1003)},
			Message: "Identifier expected.",
			Range:   f.Ranges()[1].LSRange,
		},
	})
	f.VerifyCompletions(t, "", nil)
}
