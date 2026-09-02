package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/v7/pkg/modulespecifiers"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestImportModuleSpecifierPreferenceShortest(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};
// @Filename: /project/src/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceShortest,
	})
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{""})
}

func TestImportModuleSpecifierPreferenceProjectRelative(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};
// @Filename: /project/tests/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceProjectRelative,
	})
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{""})
}

func TestImportModuleSpecifierPreferenceRelative(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};
// @Filename: /project/src/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceRelative,
	})
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{""})
}

func TestImportModuleSpecifierPreferenceProjectRelativeWithPaths(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	// Regression test: when `paths` aliases are configured but the import target
	// lives in the same package as the importing file, `project-relative` must
	// prefer the relative path rather than fall back to the `paths` alias.
	const content = `// @Filename: /project/tsconfig.json
{
  "compilerOptions": {
    "paths": {
      "@app/*": ["./src/app/*"],
      "@utils/*": ["./src/utils/*"],
    }
  }
}
// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};
// @Filename: /project/src/app/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceProjectRelative,
	})
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{""})
}

func TestImportModuleSpecifierPreferenceNonRelative(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/tsconfig.json
{
  "compilerOptions": {
    "paths": {
      "@app/*": ["./src/app/*"],
      "@utils/*": ["./src/utils/*"],
    }
  }
}
// @Filename: /project/src/utils/helper.ts
export const helperFunc = () => {};
// @Filename: /project/src/app/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierPreference:       modulespecifiers.ImportModuleSpecifierPreferenceNonRelative,
	})
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"helperFunc"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{""})
}
