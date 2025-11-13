package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixNewImportIndex(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a/index.ts
export const foo = 0;
// @Filename: /b.ts
[|/**/foo;|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/a/index.ts")
	f.GoToFile(t, "/b.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import { foo } from "./a";

foo;`,
	}, nil /*preferences*/)
}
