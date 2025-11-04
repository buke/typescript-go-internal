package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionListAtIdentifierDefinitionLocations_destructuring(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
var [x/*variable1*/
// @Filename: b.ts
var [x, y/*variable2*/
// @Filename: c.ts
var [./*variable3*/
// @Filename: d.ts
var [x, ...z/*variable4*/
// @Filename: e.ts
var {x/*variable5*/
// @Filename: f.ts
var {x, y/*variable6*/
// @Filename: g.ts
function func1({ a/*parameter1*/
// @Filename: h.ts
function func2({ a, b/*parameter2*/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, f.Markers(), nil)
}
