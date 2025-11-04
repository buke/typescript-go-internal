package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestTsxCompletionsGenericComponent(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: preserve
// @skipLibCheck: true
// @Filename: file.tsx
 declare module JSX {
     interface Element { }
     interface IntrinsicElements {
     }
     interface ElementAttributesProperty { props; }
 }

class Table<P> {
    constructor(public props: P) {}
}

type Props = { widthInCol: number; text: string; };

/**
 * @param width {number} Table width in px
 */
function createTable(width) {
    return <Table<Props> /*1*/ />
}

createTable(800);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"widthInCol",
				"text",
			},
		},
	})
}
