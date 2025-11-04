package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestReferencesForMergedDeclarations6(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Foo { }
/*1*/module /*2*/Foo {
    export interface Bar { }
    export module Bar { export interface Baz { } }
    export function Bar() { }
}

// module
import a1 = /*3*/Foo;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3")
}
