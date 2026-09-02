package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestPathCompletionsPackageJsonImportsWildcard6(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: node18
// @Filename: /package.json
{
  "name": "foo",
  "main": "dist/index.js",
  "module": "dist/index.mjs",
  "types": "dist/index.d.ts",
  "imports": {
    "#*": "./dist/*?.d.ts"
  }
}
// @Filename: /dist/index.d.ts
export const index = 0;
// @Filename: /dist/blah?.d.ts
export const blah = 0;
// @Filename: /index.mts
import { } from "/**/";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "#blah",
					Kind:   new(lsproto.CompletionItemKindFile),
					Detail: new("#blah.d.ts"),
				},
			},
		},
	})
}
