package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionOverriddenMember23(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
// @target: esnext
// @lib: esnext
const prop: symbol = Symbol();

abstract class A {
  static [prop]() {}
}

export class B extends A {
  static [|/*1*/override|] [prop]() {}
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "1")
}
