package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNoSignatureHelpOnNewKeyword(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Foo { }
new/*1*/ Foo
new /*2*/Foo(/*3*/)`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyNoSignatureHelpForMarkers(t, "1", "2")
	f.GoToMarker(t, "3")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "Foo(): Foo"})
}
