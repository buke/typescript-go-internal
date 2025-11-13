package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_order(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export const foo: number;
// @Filename: /b.ts
export const foo: number;
export const bar: number;
// @Filename: /c.ts
[|import { bar } from "./b";
foo;|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/c.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import { bar, foo } from "./b";
foo;`,
		`import { foo } from "./a";
import { bar } from "./b";
foo;`,
	}, nil /*preferences*/)
}
