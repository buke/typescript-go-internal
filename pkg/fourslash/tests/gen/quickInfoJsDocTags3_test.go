package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoJsDocTags3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: quickInfoJsDocTags3.ts
interface Foo {
    /**
     * comment
     * @author Me <me@domain.tld>
     * @see x (the parameter)
     * @param {number} x - x comment
     * @param {number} y - y comment
     * @throws {Error} comment
     */
    method(x: number, y: number): void;
}

class Bar implements Foo {
    /**/method(): void {
        throw new Error("Method not implemented.");
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
