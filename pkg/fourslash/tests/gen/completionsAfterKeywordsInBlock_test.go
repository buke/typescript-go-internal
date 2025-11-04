package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsAfterKeywordsInBlock(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class C1 {
    method(map: Map<string, string>, key: string, defaultValue: string) {
        try {
            return map.get(key)!;
        }
        catch {
            return default/*1*/
        }
    }
}
class C2 {
    method(map: Map<string, string>, key: string, defaultValue: string) {
        if (map.has(key)) {
            return map.get(key)!;
        }
        else {
            return default/*2*/
        }
    }
}
class C3 {
    method(map: Map<string, string>, key: string, returnValue: string) {
        try {
            return map.get(key)!;
        }
        catch {
            return return/*3*/
        }
    }
}
class C4 {
    method(map: Map<string, string>, key: string, returnValue: string) {
        if (map.has(key)) {
            return map.get(key)!;
        }
        else {
            return return/*4*/
        }
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, []string{"1", "2"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:    "defaultValue",
					SortText: PtrTo(string(ls.SortTextLocationPriority)),
				},
			},
		},
	})
	f.VerifyCompletions(t, []string{"3", "4"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:    "returnValue",
					SortText: PtrTo(string(ls.SortTextLocationPriority)),
				},
			},
		},
	})
}
