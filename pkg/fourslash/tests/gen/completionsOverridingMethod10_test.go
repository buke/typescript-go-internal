package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsOverridingMethod10(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.ts
// @newline: LF
interface Base {
    a: string;
    b(a: string): void;
    c(a: string): string;
    c(a: number): number;
}
class Sub implements Base {
   /*a*/
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a",
					InsertText: PtrTo("a: string;"),
					FilterText: PtrTo("a"),
					SortText:   PtrTo(string(ls.SortTextLocationPriority)),
				},
				&lsproto.CompletionItem{
					Label:      "b",
					InsertText: PtrTo("b(a: string): void {\n}"),
					FilterText: PtrTo("b"),
					SortText:   PtrTo(string(ls.SortTextLocationPriority)),
				},
				&lsproto.CompletionItem{
					Label:      "c",
					InsertText: PtrTo("c(a: string): string;\nc(a: number): number;\nc(a: unknown): string | number {\n}"),
					FilterText: PtrTo("c"),
					SortText:   PtrTo(string(ls.SortTextLocationPriority)),
				},
			},
		},
	})
}
