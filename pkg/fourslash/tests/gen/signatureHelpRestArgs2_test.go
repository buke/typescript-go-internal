package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpRestArgs2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
// @allowJs: true
// @checkJs: true
// @filename: index.js
const promisify = function (thisArg, fnName) {
    const fn = thisArg[fnName];
    return function () {
        return new Promise((resolve) => {
            fn.call(thisArg, ...arguments, /*1*/);
        });
    };
};`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineSignatureHelp(t)
}
