package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsImport_reExportDefault(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: es5
// @module: esnext
// @Filename: /a/b/impl.ts
export default function foo() {}
// @Filename: /a/index.ts
export { default as foo } from "./b/impl";
// @Filename: /use.ts
fo/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: CompletionGlobalsPlus(
				[]fourslash.CompletionsExpectedItem{
					&lsproto.CompletionItem{
						Label: "foo",
						Data: &lsproto.CompletionItemData{
							AutoImport: &lsproto.AutoImportFix{
								ModuleSpecifier: "./a",
							},
						},
						Detail:              new("(alias) function foo(): void\nexport foo"),
						Kind:                new(lsproto.CompletionItemKindFunction),
						AdditionalTextEdits: fourslash.AnyTextEdits,
						SortText:            new(string(ls.SortTextAutoImportSuggestions)),
					},
				}, false),
		},
	})
	f.VerifyApplyCodeActionFromCompletion(t, new(""), &fourslash.ApplyCodeActionFromCompletionOptions{
		Name:        "foo",
		Source:      "./a",
		Description: "Add import from \"./a\"",
		NewFileContent: new(`import { foo } from "./a";

fo`),
	})
}
