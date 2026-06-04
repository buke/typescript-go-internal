package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestParserCorruptionAfterMapInClass(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @target: esnext
// @lib: es2015
// @strict: true
class C {
    map = new Set<[|string, number|]>/*$*/

    foo() {

    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "$")
	f.Insert(t, "()")
	f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{
		{
			Code:    &lsproto.IntegerOrString{Integer: new(int32(2558))},
			Message: lsproto.StringOrMarkupContent{String: new("Expected 1 type arguments, but got 2.")},
		},
	})
}
