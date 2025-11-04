package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionUnionTypeProperty_discriminated(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type U = A | B;

interface A {
  /*aKind*/kind: "a";
  /*aProp*/prop: number;
};

interface B {
  /*bKind*/kind: "b";
  /*bProp*/prop: string;
}

const u: U = {
  [|/*kind*/kind|]: "a",
  [|/*prop*/prop|]: 0,
};
const u2: U = {
  [|/*kindBogus*/kind|]: "bogus",
  [|/*propBogus*/prop|]: 0,
};`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "kind", "prop", "kindBogus", "propBogus")
}
