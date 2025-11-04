package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoJsPropertyAssignedAfterMethodDeclaration(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noLib: true
// @allowJs: true
// @noImplicitThis: true
// @Filename: /a.js
const o = {
    test/*1*/() {
        this./*2*/test = 0;
    }
};`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "(method) test(): void", "")
	f.VerifyQuickInfoAt(t, "2", "(method) test(): void", "")
}
