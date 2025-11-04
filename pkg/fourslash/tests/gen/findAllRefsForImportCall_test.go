package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsForImportCall(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /app.ts
export function he/**/llo() {};
// @Filename: /re-export.ts
export const services = { app: setup(() => import('./app')) }
function setup<T>(importee: () => Promise<T>): T { return {} as any }
// @Filename: /indirect-use.ts
import("./re-export").then(mod => mod.services.app.hello());
// @Filename: /direct-use.ts
async function main() {
    const mod = await import("./app")
    mod.hello();
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "")
}
