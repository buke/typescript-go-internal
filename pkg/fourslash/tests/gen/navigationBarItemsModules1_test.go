package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNavigationBarItemsModules1(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare module "X.Y.Z" {}

declare module 'X2.Y2.Z2' {}

declare module "foo";

module A.B.C {
    export var x;
}

module A.B {
    export var y;
}

module A {
    export var z;
}

module A {
    module B {
        module C {
            declare var x;
        }
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
