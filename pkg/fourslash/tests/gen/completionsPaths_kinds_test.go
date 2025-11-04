package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsPaths_kinds(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /src/b.ts
not read
// @Filename: /src/dir/x.ts
not read
// @Filename: /src/a.ts
import {} from "./[|/*0*/|]";
import {} from "./[|/*1*/|]";
// @Filename: /tsconfig.json
{
    "compilerOptions": {
        "baseUrl": ".",
        "paths": {
            "foo/*": ["src/*"]
        }
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"0", "1"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "b",
					Kind:  PtrTo(lsproto.CompletionItemKindFile),
				},
				&lsproto.CompletionItem{
					Label: "dir",
					Kind:  PtrTo(lsproto.CompletionItemKindFolder),
				},
			},
		},
	})
}
