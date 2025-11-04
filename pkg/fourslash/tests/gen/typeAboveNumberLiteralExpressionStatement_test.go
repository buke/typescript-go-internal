package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestTypeAboveNumberLiteralExpressionStatement(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
// foo
1;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToBOF(t)
	f.Insert(t, "var x;\n")
}
