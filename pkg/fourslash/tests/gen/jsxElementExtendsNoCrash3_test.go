package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsxElementExtendsNoCrash3(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: index.tsx
<T extends /=>`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifySuggestionDiagnostics(t, nil)
}
