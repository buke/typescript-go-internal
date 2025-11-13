package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_quoteStyle(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export const foo: number;
// @Filename: /b.ts
[|foo;|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/b.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import { foo } from './a';

foo;`,
	}, &lsutil.UserPreferences{QuotePreference: lsutil.QuotePreference("single")})
}
