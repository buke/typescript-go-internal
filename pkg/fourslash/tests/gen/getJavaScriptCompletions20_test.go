package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetJavaScriptCompletions20(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowNonTsExtensions: true
// @Filename: file.js
/**
 * A person
 * @constructor
 * @param {string} name - The name of the person.
 * @param {number} age - The age of the person.
 */
function Person(name, age) {
    this.name = name;
    this.age = age;
}


Person.getName = 10;
Person.getNa/**/ = 10;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: CompletionFunctionMembersWithPrototypePlus(
				[]fourslash.CompletionsExpectedItem{
					"getName",
					"getNa",
					&lsproto.CompletionItem{
						Label:    "Person",
						SortText: PtrTo(string(ls.SortTextJavascriptIdentifiers)),
					},
					&lsproto.CompletionItem{
						Label:    "name",
						SortText: PtrTo(string(ls.SortTextJavascriptIdentifiers)),
					},
					&lsproto.CompletionItem{
						Label:    "age",
						SortText: PtrTo(string(ls.SortTextJavascriptIdentifiers)),
					},
				}),
		},
	})
}
