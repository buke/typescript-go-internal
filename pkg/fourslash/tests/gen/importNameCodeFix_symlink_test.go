package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_symlink(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @moduleResolution: bundler
// @noLib: true
// @Filename: /node_modules/real/index.d.ts
// @Symlink: /node_modules/link/index.d.ts
export const foo: number;
// @Filename: /a.ts
import { foo } from "link";
// @Filename: /b.ts
[|foo;|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/b.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import { foo } from "link";

foo;`,
		`import { foo } from "real";

foo;`,
	}, nil /*preferences*/)
}
