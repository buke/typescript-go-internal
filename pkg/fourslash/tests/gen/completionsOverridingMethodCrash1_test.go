package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestCompletionsOverridingMethodCrash1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @newline: LF
// @Filename: a.ts
declare class Component<T> {
    setState(stateHandler: ((oldState: T, newState: T) => void)): void;
}

class SubComponent extends Component<{}> {
    /*$*/
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "$", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "setState",
					InsertText: PtrTo("setState(stateHandler: (oldState: {}, newState: {}) => void): void {\n}"),
					FilterText: PtrTo("setState"),
					SortText:   PtrTo(string(ls.SortTextLocationPriority)),
				},
			},
		},
	})
}
