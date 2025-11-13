package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportFixes_quotePreferenceDouble_importHelpers(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @importHelpers: true
// @filename: /a.ts
export default () => {};
// @filename: /b.ts
export default () => {};
// @filename: /test.ts
import a from "./a";
[|b|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/test.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import b from "./b";
b`,
	}, nil /*preferences*/)
}
