package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNavigationItemsExportEqualsExpression(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `export = function () {}
export = function () {
    return class Foo {
    }
}

export = () => ""
export = () => {
    return class Foo {
    }
}

export = function f1() {}
export = function f2() {
    return class Foo {
    }
}

const abc = 12;
export = abc;
export = class AB {}
export = {
    a: 1,
    b: 1,
    c: {
        d: 1
    }
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
}
