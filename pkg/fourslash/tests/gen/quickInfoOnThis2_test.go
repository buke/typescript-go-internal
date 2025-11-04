package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoOnThis2(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Bar<T> {
    public explicitThis(this: this) {
        console.log(th/*1*/is);
    }
    public explicitClass(this: Bar<T>) {
        console.log(thi/*2*/s);
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "this: this", "")
	f.VerifyQuickInfoAt(t, "2", "this: Bar<T>", "")
}
