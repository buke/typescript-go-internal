package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestOrganizeImports4(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import * as something from "path";/** 
 * some comment here
 * and there
 */
import * as somethingElse from "anotherpath";
import * as AnotherThing from "somepath";/** 
 * some comment here
 * and there
 */
import * as AnotherThingElse from "someotherpath";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		``,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
