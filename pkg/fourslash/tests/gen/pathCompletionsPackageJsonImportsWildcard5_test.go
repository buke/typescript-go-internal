package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestPathCompletionsPackageJsonImportsWildcard5(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: node18
// @Filename: /package.json
{
  "name": "foo",
  "main": "dist/index.js",
  "module": "dist/index.mjs",
  "types": "dist/index.d.ts",
  "imports": {
    "#*": {
      "import": {
        "types": "./dist/types/*.d.mts",
        "default": "./dist/esm/*.mjs"
      },
      "default": {
        "types": "./dist/types/*.d.ts",
        "default": "./dist/cjs/*.js"
      }
    },
    "#only-in-cjs": {
      "require": {
        "types": "./dist/types/only-in-cjs/index.d.ts",
        "default": "./dist/cjs/only-in-cjs/index.js"
      }
    }
  }
}
// @Filename: /dist/types/index.d.mts
export const index = 0;
// @Filename: /dist/types/index.d.ts
export const index = 0;
// @Filename: /dist/types/blah.d.mts
export const blah = 0;
// @Filename: /dist/types/blah.d.ts
export const blah = 0;
// @Filename: /dist/types/only-in-cjs/index.d.ts
export const onlyInCjs = 0;
// @Filename: /index.mts
import { } from "/**/";`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "#blah",
					Kind:  PtrTo(lsproto.CompletionItemKindFile),
				},
				&lsproto.CompletionItem{
					Label: "#index",
					Kind:  PtrTo(lsproto.CompletionItemKindFile),
				},
			},
		},
	})
}
