package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionImportedNames6(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: b.ts
import [|/*moduleAliasDefinition*/alias|] = require("./a");
// @Filename: a.ts
/*moduleDefinition*/export module Module {
}
export class Class {
    private f;
}
export interface Interface {
    x;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "moduleAliasDefinition")
}
