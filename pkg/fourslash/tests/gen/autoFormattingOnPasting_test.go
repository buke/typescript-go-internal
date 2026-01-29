package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestAutoFormattingOnPasting(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `module TestModule {
/**/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.Paste(t, " class TestClass{\nprivate   foo;\npublic testMethod( )\n{}\n}")
	f.VerifyCurrentFileContent(t, `module TestModule {
    class TestClass {
        private foo;
        public testMethod() { }
    }
}`)
}
