package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsImportYieldExpression(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export function a() {}
// @Filename: /b.ts
function *f() {
  yield a/**/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyApplyCodeActionFromCompletion(t, new(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:        "a",
		Source:      "./a",
		Description: "Add import from \"./a\"",
		NewFileContent: new(`import { a } from "./a";

function *f() {
  yield a
}`),
	})
}
