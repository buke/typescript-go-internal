package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestDocumentHighlightDefaultInSwitch(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const foo = 'foo';
[|switch|] (foo) {
   [|case|] 'foo':
       [|break|];
   [|default|]:
       [|break|];
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, f.Ranges()[1], f.Ranges()[4])
}
