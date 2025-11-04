package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestAutoImportTypeOnlyPreferred1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @verbatimModuleSyntax: true
// @module: esnext
// @moduleResolution: bundler
// @Filename: /ts.d.ts
declare namespace ts {
  interface SourceFile {
      text: string;
  }
  function createSourceFile(): SourceFile;
}
export = ts;
// @Filename: /types.ts
export interface VFS {
  getSourceFile(path: string): ts/**/
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "ts",
					Data: PtrTo(any(&ls.CompletionItemData{
						AutoImport: &ls.AutoImportData{
							ModuleSpecifier: "./ts",
						},
					})),
					AdditionalTextEdits: fourslash.AnyTextEdits,
					SortText:            PtrTo(string(ls.SortTextAutoImportSuggestions)),
				},
			},
		},
	}).AndApplyCodeAction(t, &fourslash.CompletionsExpectedCodeAction{
		Name:        "ts",
		Source:      "./ts",
		Description: "Add import from \"./ts\"",
		NewFileContent: `import type ts from "./ts";

export interface VFS {
  getSourceFile(path: string): ts
}`,
	})
}
