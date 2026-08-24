package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/testutil"
	"github.com/buke/typescript-go-internal/pkg/testutil/contentmappertest"
)

func TestContentMapperSynthesizedDocumentSymbols(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := newContentMapperFourslash(t, `// @Filename: /app.vue
component source with no direct TypeScript span/**/
`, contentmappertest.SynthesizingMapper, ".vue")
	defer done()

	f.GoToMarker(t, "")
	f.VerifyBaselineDocumentSymbol(t)
}

func TestContentMapperSupplementalDocumentSymbols(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := newContentMapperFourslash(t, `// @Filename: /app.astro
export function supplementalSymbol() {}
`, contentmappertest.SupplementalMapper, ".astro")
	defer done()

	f.GoToFile(t, "/app.astro")
	f.VerifyBaselineDocumentSymbol(t)
}
