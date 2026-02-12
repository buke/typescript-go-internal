package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestOrganizeImportsPathsUnicode3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import * as B from "./B";
import * as À from "./À";
import * as A from "./A";

console.log(A, À, B);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import * as À from "./À";
import * as A from "./A";
import * as B from "./B";

console.log(A, À, B);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:      core.TSFalse,
			OrganizeImportsCollation:       lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsAccentCollation: false,
		},
	)
	f.VerifyOrganizeImports(t,
		`import * as A from "./A";
import * as À from "./À";
import * as B from "./B";

console.log(A, À, B);`,
		lsproto.CodeActionKindSourceOrganizeImports,
		&lsutil.UserPreferences{
			OrganizeImportsIgnoreCase:      core.TSFalse,
			OrganizeImportsCollation:       lsutil.OrganizeImportsCollationUnicode,
			OrganizeImportsAccentCollation: true,
		},
	)
}
