package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestInlayHintsNoHintWhenArgumentMatchesName(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo (a: number, b: number) {}
declare const a: 1;
foo(a, 2);
declare const v: any;
foo(v.a, v.a);
foo(v.b, v.b);
foo(v.c, v.c);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{IncludeInlayParameterNameHints: lsutil.IncludeInlayParameterNameHintsAll, IncludeInlayParameterNameHintsWhenArgumentMatchesName: false})
}
