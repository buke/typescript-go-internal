package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionDifferentFile(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: goToDefinitionDifferentFile_Definition.ts
var /*remoteVariableDefinition*/remoteVariable;
function /*remoteFunctionDefinition*/remoteFunction() { }
class /*remoteClassDefinition*/remoteClass { }
interface /*remoteInterfaceDefinition*/remoteInterface{ }
module /*remoteModuleDefinition*/remoteModule{ export var foo = 1;}
// @Filename: goToDefinitionDifferentFile_Consumption.ts
/*remoteVariableReference*/remoteVariable = 1;
/*remoteFunctionReference*/remoteFunction();
var foo = new /*remoteClassReference*/remoteClass();
class fooCls implements /*remoteInterfaceReference*/remoteInterface { }
var fooVar = /*remoteModuleReference*/remoteModule.foo;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, false, "remoteVariableReference", "remoteFunctionReference", "remoteClassReference", "remoteInterfaceReference", "remoteModuleReference")
}
