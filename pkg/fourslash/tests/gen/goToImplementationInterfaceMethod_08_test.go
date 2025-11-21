package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToImplementationInterfaceMethod_08(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Foo {
    hello (): void;
}

class SuperBar implements Foo {
   [|hello|]() {}
}

class Bar extends SuperBar {
   whatever() { this.he/*function_call*/llo(); }
}

class SubBar extends Bar {
   [|hello|]() {}
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToImplementation(t, "function_call")
}
