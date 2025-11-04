package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionOverriddenMember20(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
// @target: esnext
// @lib: esnext
const prop = "foo" as const;

abstract class A {
  readonly /*2*/[prop] = "A";
}

export class B extends A {
  [|/*1*/override|] readonly [prop] = "B";
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "1")
}
