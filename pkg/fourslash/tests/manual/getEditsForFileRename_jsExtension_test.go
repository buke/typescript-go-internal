package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetEditsForFileRename_jsExtension(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: /src/a.js
export const a = 0;
// @Filename: /b.js
import { a } from "./src/a.js";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWillRenameFilesEdits(t, "/b.js", "/src/b.js", map[string]string{
		"/src/b.js": `import { a } from "./a.js";`,
	}, nil /*preferences*/)
}
