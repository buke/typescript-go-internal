package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoSignatureOptionalParameterFromUnion1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare const optionals:
  | ((a?: { a: true }) => unknown)
  | ((b?: { b: true }) => unknown);

/**/optionals();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "const optionals: (arg0?: {\n    a: true;\n} & {\n    b: true;\n}) => unknown", "")
}
