package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoOnFunctionPropertyReturnedFromGenericFunction3(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function createProps<T>(t: T) {
  const getProps = () => {}
  const createVariants = () => {}

  getProps.createVariants = createVariants;
  return getProps;
}

createProps({})./**/createVariants();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "(property) getProps<{}>.createVariants: () => void", "")
}
