package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFormatNestedClassWithOpenBraceOnNewLines(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module A
{
    class B {
        /*1*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	opts168 := f.GetOptions()
	opts168.FormatCodeSettings.PlaceOpenBraceOnNewLineForControlBlocks = core.TSTrue
	f.Configure(t, opts168)
	opts232 := f.GetOptions()
	opts232.FormatCodeSettings.PlaceOpenBraceOnNewLineForFunctions = core.TSTrue
	f.Configure(t, opts232)
	f.GoToMarker(t, "1")
	f.Insert(t, "}")
	f.VerifyCurrentFileContent(t, `module A
{
    class B
    {
    }
}`)
}
