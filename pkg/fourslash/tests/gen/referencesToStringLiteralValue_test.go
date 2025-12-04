package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestReferencesToStringLiteralValue(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `const s: string = "some /*1*/ string";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.MarkTestAsStradaServer()
	f.VerifyBaselineFindAllReferences(t, "1")
}
