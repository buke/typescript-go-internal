package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestOrganizeImports6(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import * as something from "path"; /* small comment */ // single line one.
/* some comment here
* and there
*/
import * as somethingElse from "anotherpath";
import * as anotherThing from "someopath"; /* small comment */ // single line one.
/* some comment here
* and there
*/
import * as anotherThingElse from "someotherpath";

anotherThing;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`/* some comment here
* and there
*/
import * as anotherThing from "someopath"; /* small comment */ // single line one.
/* some comment here
* and there
*/

anotherThing;`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
