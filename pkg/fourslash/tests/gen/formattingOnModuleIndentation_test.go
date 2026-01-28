package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFormattingOnModuleIndentation(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `  module     Foo    {
    export    module    A  .   B  .   C     {      }/**/
               }`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToBOF(t)
	f.VerifyCurrentLineContent(t, `module Foo {`)
	f.GoToMarker(t, "")
	f.VerifyCurrentLineContent(t, `    export module A.B.C { }`)
	f.GoToEOF(t)
	f.VerifyCurrentLineContent(t, `}`)
}
