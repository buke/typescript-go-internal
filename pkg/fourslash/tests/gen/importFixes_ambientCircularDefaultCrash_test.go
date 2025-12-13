package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportFixes_ambientCircularDefaultCrash(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /home/src/workspaces/project/tsconfig.json
{
  "compilerOptions": {
    "module": "preserve"
  }
}
// @Filename: /home/src/workspaces/project/types.d.ts
declare module "mymod" {
  import mymod from "mymod";
  export default mymod;
}
// @Filename: /home/src/workspaces/project/index.ts
my/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.MarkTestAsStradaServer()
	f.VerifyImportFixModuleSpecifiers(t, "", []string{}, nil /*preferences*/)
}
