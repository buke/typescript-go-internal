package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixDefaultExport5(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @moduleResolution: bundler
// @Filename: /node_modules/hooks/useFoo.ts
declare const _default: () => void;
export default _default;
// @Filename: /test.ts
[|useFoo|];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{AutoImportEntrypointDirectorySearch: core.TSTrue})
	f.GoToFile(t, "/test.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`import useFoo from "hooks/useFoo";

useFoo`,
	}, nil /*preferences*/)
}
