package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionMultipleDefinitions(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
interface /*interfaceDefinition1*/IFoo {
    instance1: number;
}
// @Filename: b.ts
interface /*interfaceDefinition2*/IFoo {
    instance2: number;
}

interface /*interfaceDefinition3*/IFoo {
    instance3: number;
}

var ifoo: [|IFo/*interfaceReference*/o|];
// @Filename: c.ts
module /*moduleDefinition1*/Module {
    export class c1 { }
}
// @Filename: d.ts
module /*moduleDefinition2*/Module {
    export class c2 { }
}
// @Filename: e.ts
[|Modul/*moduleReference*/e|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "interfaceReference", "moduleReference")
}
