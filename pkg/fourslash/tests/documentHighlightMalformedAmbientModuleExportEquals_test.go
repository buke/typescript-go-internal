package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestDocumentHighlightMalformedAmbientModuleExportEquals(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `// @Filename: /a.d.ts
declare moduleu "m" {
  interface A { x: 1 }
  function f(): A[];
  /*m*/export = f;
}`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, "m")
}
