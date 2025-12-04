package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestTsxRename3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
declare module JSX {
    interface Element { }
    interface IntrinsicElements {
    }
    interface ElementAttributesProperty { props }
}
class MyClass {
  props: {
    [|[|{| "contextRangeIndex": 0 |}name|]?: string;|]
    size?: number;
}


var x = <MyClass [|[|{| "contextRangeIndex": 2 |}name|]='hello'|]/>;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRenameAtRangesWithText(t, nil /*preferences*/, "name")
}
