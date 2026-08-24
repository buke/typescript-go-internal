package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/testutil"
	"github.com/buke/typescript-go-internal/pkg/testutil/contentmappertest"
)

func TestContentMapperDiagnosticCodeDoesNotSelectTypeScriptFix(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := newContentMapperFourslash(t, `// @Filename: /app.box
export function f/*query*/oo() { return 1; }
`, contentmappertest.DiagnosticCodeCollisionMapper, ".box")
	defer done()

	f.GoToMarker(t, "query")
	f.VerifyCodeFixNotAvailable(t)
}
