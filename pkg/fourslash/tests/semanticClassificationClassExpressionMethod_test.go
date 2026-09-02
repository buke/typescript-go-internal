package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestSemanticClassificationClassExpressionMethod(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var x = class C {
  equals(other: C) { return this == other; }
};`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifySemanticTokens(t, []fourslash.SemanticToken{
		{Type: "class.declaration", Text: "x"},
		{Type: "class.declaration", Text: "C"},
		{Type: "method.declaration", Text: "equals"},
		{Type: "parameter.declaration", Text: "other"},
		{Type: "class", Text: "C"},
		{Type: "parameter", Text: "other"},
	})
}
