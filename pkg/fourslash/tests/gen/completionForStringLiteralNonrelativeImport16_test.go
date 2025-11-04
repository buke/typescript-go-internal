package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionForStringLiteralNonrelativeImport16(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: tsconfig.json
{
    "compilerOptions": {
        "resolveJsonModule": false,
        "baseUrl": "./",
        "paths": {
            "module1/path1": ["some/path/whatever.ts"],
        }
    }
}
// @Filename: test0.ts
import * as foo1 from "m/*first*/
import * as foo1 from "module1/pa/*second*/
// @Filename: some/path/whatever.ts
export var x = 9;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"first"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"test0",
				"some",
				&lsproto.CompletionItem{
					Label: "module1/path1",
				},
			},
		},
	})
	f.VerifyCompletions(t, []string{"second"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "module1/path1",
				},
			},
		},
	})
}
