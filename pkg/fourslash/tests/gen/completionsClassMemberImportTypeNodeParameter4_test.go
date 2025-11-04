package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsClassMemberImportTypeNodeParameter4(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: node18
// @FileName: /other/cls.d.ts
export declare class Cls {
  method(
    param: import("./doesntexist.js").Foo,
  ): import("./doesntexist.js").Foo;
}
// @FileName: /index.d.ts
import { Cls } from "./other/cls.js";

export declare class Derived extends Cls {
  /*1*/
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:               "method",
					InsertText:          PtrTo("method(param: import(\"./doesntexist.js\").Foo);"),
					FilterText:          PtrTo("method"),
					AdditionalTextEdits: fourslash.AnyTextEdits,
				},
			},
		},
	})
}
