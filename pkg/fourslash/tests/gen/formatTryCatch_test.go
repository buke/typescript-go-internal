package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFormatTryCatch(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function test() {
    /*try*/try {
    }
    /*catch*/catch (e) {
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.FormatDocument(t, "")
	f.FormatDocument(t, "")
	f.GoToMarker(t, "try")
	f.VerifyCurrentLineContent(t, `    try {`)
	f.GoToMarker(t, "catch")
	f.VerifyCurrentLineContent(t, `    catch (e) {`)
}
