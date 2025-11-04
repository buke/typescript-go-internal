package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllReferencesNonExistentExportBinding(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /tsconfig.json
 { "compilerOptions": { "module": "commonjs" } }
// @filename: /bar.ts
import { Foo/**/ } from "./foo";
// @filename: /foo.ts
export { Foo }`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "")
}
