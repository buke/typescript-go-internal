package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsFileImportNoTypes2(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowJs: true
// @Filename: /default.ts
export default class TestDefaultClass {}
// @Filename: /defaultType.ts
export default interface TestDefaultInterface {}
// @Filename: /reExport/toReExport.ts
export class TestClassReExport {}
export interface TestInterfaceReExport {}
// @Filename: /reExport/index.ts
export { TestClassReExport, TestInterfaceReExport } from './toReExport';
// @Filename: /exportList.ts
class TestClassExportList {};
interface TestInterfaceExportList {};
export { TestClassExportList, TestInterfaceExportList };
// @Filename: /baseline.ts
export class TestClassBaseline {}
export interface TestInterfaceBaseline {}
// @Filename: /a.js
import /**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "TestClassBaseline",
					InsertText: PtrTo("import { TestClassBaseline } from \"./baseline\";"),
					Data: PtrTo(any(&ls.CompletionItemData{
						AutoImport: &ls.AutoImportData{
							ModuleSpecifier: "./baseline",
						},
					})),
				},
				&lsproto.CompletionItem{
					Label:      "TestClassExportList",
					InsertText: PtrTo("import { TestClassExportList } from \"./exportList\";"),
					Data: PtrTo(any(&ls.CompletionItemData{
						AutoImport: &ls.AutoImportData{
							ModuleSpecifier: "./exportList",
						},
					})),
				},
				&lsproto.CompletionItem{
					Label:      "TestClassReExport",
					InsertText: PtrTo("import { TestClassReExport } from \"./reExport\";"),
					Data: PtrTo(any(&ls.CompletionItemData{
						AutoImport: &ls.AutoImportData{
							ModuleSpecifier: "./reExport",
						},
					})),
				},
				&lsproto.CompletionItem{
					Label:      "TestDefaultClass",
					InsertText: PtrTo("import TestDefaultClass from \"./default\";"),
					Data: PtrTo(any(&ls.CompletionItemData{
						AutoImport: &ls.AutoImportData{
							ModuleSpecifier: "./default",
						},
					})),
				},
			},
		},
	})
}
