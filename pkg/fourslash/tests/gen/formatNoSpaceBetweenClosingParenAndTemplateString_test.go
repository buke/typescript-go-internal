package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFormatNoSpaceBetweenClosingParenAndTemplateString(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `foo() ` + "`" + `abc` + "`" + `;
bar()` + "`" + `def` + "`" + `;
baz()` + "`" + `a${x}b` + "`" + `;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, "foo()`abc`;\nbar()`def`;\nbaz()`a${x}b`;")
}
