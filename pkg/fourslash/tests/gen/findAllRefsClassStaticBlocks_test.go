package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsClassStaticBlocks(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class ClassStaticBocks {
    static x;
    [|[|/*classStaticBocks1*/static|] {}|]
    static y;
    [|[|/*classStaticBocks2*/static|] {}|]
    static y;
    [|[|/*classStaticBocks3*/static|] {}|]
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "classStaticBocks1", "classStaticBocks2", "classStaticBocks3")
}
