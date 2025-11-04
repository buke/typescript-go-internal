package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsForImportCallType(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /app.ts
export function he/**/llo() {};
// @Filename: /re-export.ts
export type app = typeof import("./app")
// @Filename: /indirect-use.ts
import type { app } from "./re-export";
declare const app: app
app.hello();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "")
}
