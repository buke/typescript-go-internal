package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestClassMemberCompletionKeepsNameFallback(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class B {
	constructor(public value: string) {}
}
class C extends B {
	/*a*/
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCompletions(t, "a", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "value",
					InsertText: new("value"),
				},
			},
		},
		UserPreferences: &lsutil.UserPreferences{IncludeCompletionsWithClassMemberSnippets: core.TSTrue},
	})
}

func TestImplementClassFixDoesNotAddInvalidOverride(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noImplicitOverride: true
class B {
    method() {}
}
class C implements B {[| |]}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCodeFix(t, fourslash.VerifyCodeFixOptions{
		Description: "Implement interface 'B'",
		NewFileContent: `class B {
    method() {}
}
class C implements B {
    method(): void {
        throw new Error("Method not implemented.");
    }
}`,
		Index: 0,
	})
}
