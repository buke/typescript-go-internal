package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoFromContextualUnionType2(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
function test1(arg: { prop: "foo" }) {}
test1({ /*1*/prop: "bar" });

function test2(arg: { prop: "foo" } | undefined) {}
test2({ /*2*/prop: "bar" });`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "(property) prop: \"foo\"", "")
	f.VerifyQuickInfoAt(t, "2", "(property) prop: \"foo\"", "")
}
