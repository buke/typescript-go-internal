package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/v7/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/v7/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestAutoImportDefaultPascalCase(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @module: esnext
// @moduleResolution: bundler

// @Filename: /src/components/ChargerHeader.tsx
export default function ChargerHeader() {
  return null;
}

// @Filename: /src/screens/SomeScreen.tsx
export function SomeScreen() {
  return <ChargerHeader/*1*/
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
		},
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"ChargerHeader"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{"1"})
}

func TestAutoImportDefaultPascalCaseAnonymous(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @module: esnext
// @moduleResolution: bundler

// @Filename: /src/components/ChargerHeader.tsx
export default function() {
  return null;
}

// @Filename: /src/screens/SomeScreen.tsx
export function SomeScreen() {
  return <ChargerHeader/*1*/
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
		},
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"ChargerHeader"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{"1"})
}

func TestAutoImportDefaultPascalCaseCaseInsensitive(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @module: esnext
// @moduleResolution: bundler
// @useCaseSensitiveFileNames: false

// @Filename: /src/components/ChargerHeader.tsx
export default function ChargerHeader() {
  return null;
}

// @Filename: /src/screens/SomeScreen.tsx
export function SomeScreen() {
  return <ChargerHeader/*1*/
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
		},
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"ChargerHeader"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{"1"})
}

func TestAutoImportDefaultPascalCaseAnonymousCaseInsensitive(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @module: esnext
// @moduleResolution: bundler
// @useCaseSensitiveFileNames: false

// @Filename: /src/components/ChargerHeader.tsx
export default function() {
  return null;
}

// @Filename: /src/screens/SomeScreen.tsx
export function SomeScreen() {
  return <ChargerHeader/*1*/
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
		},
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"ChargerHeader"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{"1"})
}

func TestAutoImportDefaultPascalCaseReexportCaseInsensitive(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @module: esnext
// @moduleResolution: bundler
// @useCaseSensitiveFileNames: false

// @Filename: /src/components/ChargerHeader.tsx
export default function() {
  return null;
}

// @Filename: /src/components/index.ts
export { default } from "./ChargerHeader";

// @Filename: /src/screens/SomeScreen.tsx
export function SomeScreen() {
  return <ChargerHeader/*1*/
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
		},
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"ChargerHeader"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{"1"})
}

func TestAutoImportDefaultPascalCaseAliasCaseInsensitive(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @module: esnext
// @moduleResolution: bundler
// @useCaseSensitiveFileNames: false

// @Filename: /src/components/ChargerHeader.tsx
function ChargerHeader() {
  return null;
}
export default ChargerHeader;

// @Filename: /src/screens/SomeScreen.tsx
export function SomeScreen() {
  return <ChargerHeader/*1*/
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		UserPreferences: &lsutil.UserPreferences{
			IncludeCompletionsForModuleExports:    core.TSTrue,
			IncludeCompletionsForImportStatements: core.TSTrue,
		},
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"ChargerHeader"},
		},
	})
	f.BaselineAutoImportsCompletions(t, []string{"1"})
}
