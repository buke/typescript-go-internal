package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesAfterEdit(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*0*/
interface A {
    foo: string;
}
function foo(x: A) {
    x.f/*1*/oo
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, "1")
	f.GoToMarker(t, "0")
	f.Insert(t, "\n")
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, "1")
}
