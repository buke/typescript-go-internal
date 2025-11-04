package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesSetAndGet2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Foo {
    set bar(b: any) {
    }

    public get bar(): any {
        return undefined;
    }

    public [|set|] set(s: any) {
    }

    public [|get|] set(): any {
        return undefined;
    }

    public set get(g: any) {
    }

    public get get(): any {
        return undefined;
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
