package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsImportModuleAugmentationWithJS(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @checkJs: true
// @noEmit: true
// @Filename: /test.js
class Abcde {
    x
}

module.exports = {
    Abcde
};
// @Filename: /index.ts
export {};
declare module "./test" {
    interface Abcde { b: string }
}

Abcde/**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyApplyCodeActionFromCompletion(t, PtrTo(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:        "Abcde",
		Source:      "./test",
		Description: "Add import from \"./test\"",
		NewFileContent: PtrTo(`import { Abcde } from "./test";

export {};
declare module "./test" {
    interface Abcde { b: string }
}

Abcde`),
	})
}
