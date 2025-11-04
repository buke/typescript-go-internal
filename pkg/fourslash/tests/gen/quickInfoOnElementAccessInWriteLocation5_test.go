package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoOnElementAccessInWriteLocation5(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
interface Serializer {
  set value(v: string | number);
  get value(): string;
}
declare let box: Serializer;
box['value'/*1*/] += 10;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "1", "(property) Serializer.value: string | number", "")
}
