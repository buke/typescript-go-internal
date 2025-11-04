package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestNgProxy2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: tsconfig.json
{
    "compilerOptions": {
        "plugins": [
            { "name": "invalidmodulename" }
        ]
    },
    "files": ["a.ts"]
}
// @Filename: a.ts
let x = [1, 2];
x/**/
`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifyQuickInfoIs(t, "let x: number[]", "")
}
