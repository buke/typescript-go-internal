package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionListAtIdentifierDefinitionLocations_Generics(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface A</*genericName1*/
class A</*genericName2*/
class B<T, /*genericName3*/
class A{
     f</*genericName4*/
function A</*genericName5*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, f.Markers(), nil)
}
