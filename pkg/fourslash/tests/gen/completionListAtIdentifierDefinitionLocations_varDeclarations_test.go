package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionListAtIdentifierDefinitionLocations_varDeclarations(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var aa = 1;
var /*varName1*/
var a/*varName2*/
var a2,/*varName3*/
var a2, a/*varName4*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, f.Markers(), nil)
}
