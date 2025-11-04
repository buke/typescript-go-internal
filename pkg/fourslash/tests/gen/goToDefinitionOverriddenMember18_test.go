package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionOverriddenMember18(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
// @target: esnext
// @lib: esnext
const entityKind = Symbol.for("drizzle:entityKind");

abstract class MySqlColumn {
  readonly /*2*/[entityKind]: string = "MySqlColumn";
}

export class MySqlVarBinary extends MySqlColumn {
  [|/*1*/override|] readonly [entityKind]: string = "MySqlVarBinary";
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "1")
}
