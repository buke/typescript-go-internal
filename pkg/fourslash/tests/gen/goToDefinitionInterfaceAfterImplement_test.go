package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionInterfaceAfterImplement(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface /*interfaceDefinition*/sInt {
    sVar: number;
    sFn: () => void;
}

class iClass implements /*interfaceReference*/sInt {
    public sVar = 1;
    public sFn() {
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, false, "interfaceReference")
}
