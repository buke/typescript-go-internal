package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestReferencesInEmptyFileWithMultipleProjects(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /home/src/workspaces/project/a/tsconfig.json
{ "files": ["a.ts"] }
// @Filename: /home/src/workspaces/project/a/a.ts
/// <reference path="../b/b.ts" />
/*1*/;
// @Filename: /home/src/workspaces/project/b/tsconfig.json
{ "files": ["b.ts"] }
// @Filename: /home/src/workspaces/project/b/b.ts
/*2*/;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2")
}
