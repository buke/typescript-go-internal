package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoEnumMembersAcceptNonAsciiStrings(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `enum Demo {
    /*Emoji*/Emoji = '🍎',
    /*Hebrew*/Hebrew = 'תפוח',
    /*Chinese*/Chinese = '苹果',
    /*Japanese*/Japanese = 'りんご',
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "Emoji", "(enum member) Demo.Emoji = \"🍎\"", "")
	f.VerifyQuickInfoAt(t, "Hebrew", "(enum member) Demo.Hebrew = \"תפוח\"", "")
	f.VerifyQuickInfoAt(t, "Chinese", "(enum member) Demo.Chinese = \"苹果\"", "")
	f.VerifyQuickInfoAt(t, "Japanese", "(enum member) Demo.Japanese = \"りんご\"", "")
}
