package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/modulespecifiers"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportModuleSpecifierEndingAuto(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/helper/index.ts
export const helperFunc = () => {};
// @Filename: /project/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierEnding:           modulespecifiers.ImportModuleSpecifierEndingPreferenceAuto,
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

func TestImportModuleSpecifierEndingMinimal(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/helper/index.ts
export const helperFunc = () => {};
// @Filename: /project/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierEnding:           modulespecifiers.ImportModuleSpecifierEndingPreferenceMinimal,
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

func TestImportModuleSpecifierEndingIndex(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/helper/index.ts
export const helperFunc = () => {};
// @Filename: /project/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierEnding:           modulespecifiers.ImportModuleSpecifierEndingPreferenceIndex,
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

func TestImportModuleSpecifierEndingJs(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /project/helper/index.ts
export const helperFunc = () => {};
// @Filename: /project/index.ts
helper/**/`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{
		IncludeCompletionsForModuleExports:    core.TSTrue,
		IncludeCompletionsForImportStatements: core.TSTrue,
		ImportModuleSpecifierEnding:           modulespecifiers.ImportModuleSpecifierEndingPreferenceJs,
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
