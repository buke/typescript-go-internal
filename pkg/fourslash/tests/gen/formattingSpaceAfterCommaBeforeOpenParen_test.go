package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFormattingSpaceAfterCommaBeforeOpenParen(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `foo(a,(b))/*1*/
foo(a,(<b>c).d)/*2*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.Insert(t, ";")
	f.VerifyCurrentLineContent(t, `foo(a, (b));`)
	f.GoToMarker(t, "2")
	f.Insert(t, ";")
	f.VerifyCurrentLineContent(t, `foo(a, (<b>c).d);`)
}
