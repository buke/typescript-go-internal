package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestGetEditsForFileRename_resolveJsonModule(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @resolveJsonModule: true
// @Filename: /a.ts
import text from "./message.json";
// @Filename: /message.json
{}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWillRenameFilesEdits(t, "/a.ts", "/src/a.ts", map[string]string{
		"/src/a.ts": `import text from "../message.json";`,
	}, nil /*preferences*/)
}
