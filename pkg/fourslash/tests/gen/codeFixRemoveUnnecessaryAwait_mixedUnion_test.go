package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCodeFixRemoveUnnecessaryAwait_mixedUnion(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @target: esnext
async function fn1(a: Promise<void> | void) {
  await a;
}

async function fn2<T extends Promise<void> | void>(a: T) {
  await a;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifySuggestionDiagnostics(t, nil)
}
