package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/ls"
	"github.com/buke/typescript-go-internal/v7/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestCompletionFilterText3(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
declare const foo1: { b: number; "a bc": string; };
if (true) {
    foo1[|.|]/*1*/
} 
else {
    foo1[|.a|]/*2*/
}

declare const foo2: { b: number; "a bc": string; } | undefined;
if (true) {
    foo2[|.|]/*3*/
} else if (false) {
    foo2[|.a|]/*4*/
} else {
    foo2[|?.|]/*5*/
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "[\"a bc\"]",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "[\"a bc\"]",
							Range:   f.Ranges()[1].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "3", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("?.[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "?.[\"a bc\"]",
							Range:   f.Ranges()[2].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "4", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("?.[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "?.[\"a bc\"]",
							Range:   f.Ranges()[3].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "5", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("?.[\"a bc\"]"),
					FilterText: new("?.a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "?.[\"a bc\"]",
							Range:   f.Ranges()[4].LSRange,
						},
					},
				},
			},
		},
	})
}
