package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_jsxOpeningTagImportDefault(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: commonjs
// @jsx: react-jsx
// @Filename: /component.tsx
export default function (props: any) {}
// @Filename: /index.tsx
export function Index() {
    return <Component/**/ />;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifyImportFixAtPosition(t, []string{
		`import Component from "./component";

export function Index() {
    return <Component />;
}`,
	}, nil /*preferences*/)
}
