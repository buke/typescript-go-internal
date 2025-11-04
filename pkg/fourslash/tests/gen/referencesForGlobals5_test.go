package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestReferencesForGlobals5(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: referencesForGlobals_1.ts
module globalModule {
    export var x;
}

/*1*/import /*2*/globalAlias = globalModule;
// @Filename: referencesForGlobals_2.ts
var m = /*3*/globalAlias;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3")
}
