package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionDecoratorOverloads(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Target: ES6
// @experimentaldecorators: true
async function f() {}

function /*defDecString*/dec(target: any, propertyKey: string): void;
function /*defDecSymbol*/dec(target: any, propertyKey: symbol): void;
function dec(target: any, propertyKey: string | symbol) {}

declare const s: symbol;
class C {
    @[|/*useDecString*/dec|] f() {}
    @[|/*useDecSymbol*/dec|] [s]() {}
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "useDecString", "useDecSymbol")
}
