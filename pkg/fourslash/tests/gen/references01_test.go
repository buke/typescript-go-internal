package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestReferences01(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /home/src/workspaces/project/referencesForGlobals_1.ts
class /*0*/globalClass {
    public f() { }
}
// @Filename: /home/src/workspaces/project/referencesForGlobals_2.ts
///<reference path="referencesForGlobals_1.ts" />
var c = /*1*/globalClass();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.MarkTestAsStradaServer()
	f.VerifyBaselineFindAllReferences(t, "1")
}
