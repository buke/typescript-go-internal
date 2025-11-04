package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestReferencesForClassMembersExtendingGenericClass(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class Base<T> {
    /*a1*/a: this;
    /*method1*/method<U>(a?:T, b?:U): this { }
}
class MyClass extends Base<number> {
    /*a2*/a;
    /*method2*/method() { }
}

var c: MyClass;
c./*a3*/a;
c./*method3*/method();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "a1", "a2", "a3", "method1", "method2", "method3")
}
