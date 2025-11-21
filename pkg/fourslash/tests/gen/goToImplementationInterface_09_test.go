package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToImplementationInterface_09(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: def.d.ts
export interface Interface { P: number }
// @Filename: ref.ts
import { Interface } from "./def";
const c: I/*ref*/nterface = [|{ P: 2 }|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToImplementation(t, "ref")
}
