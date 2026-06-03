package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestOrganizeImports_removeUnused_preservesMultiline(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    a,
    b,
    c,
} from "module";

export { a, b, c };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(
		t,
		`import {
    a,
    b,
    c,
} from "module";

export { a, b, c };`,
		lsproto.CodeActionKindSourceRemoveUnusedImports,
		nil,
	)
}

func TestOrganizeImports_removeUnused_preservesMultilineWithRemoval(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import {
    a,
    b,
    c,
} from "module";

export { a, c };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(
		t,
		`import {
    a,
    c
} from "module";

export { a, c };`,
		lsproto.CodeActionKindSourceRemoveUnusedImports,
		nil,
	)
}
