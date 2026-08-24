package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsOverridingMethod1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @newline: LF
// @Filename: h.ts
// @noImplicitOverride: true
class HBase {
    foo(a: string): void {}
}

class HSub extends HBase {
    [|f/*h*/|]
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "h", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "foo",
					TextEdit:   InsertReplaceTextEdit("override foo(a: string): void {\n}", f.Ranges()[0].LSRange),
					FilterText: new("foo"),
					SortText:   new(string(ls.SortTextLocationPriority)),
				},
			},
		},
		UserPreferences: &lsutil.UserPreferences{IncludeCompletionsWithClassMemberSnippets: core.TSTrue},
	})
}
