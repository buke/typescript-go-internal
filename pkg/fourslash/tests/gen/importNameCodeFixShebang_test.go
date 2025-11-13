package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFixShebang(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.ts
export const foo = 0;
// @Filename: /b.ts
[|#!/usr/bin/env node
foo/**/|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/a.ts")
	f.GoToFile(t, "/b.ts")
	f.VerifyImportFixAtPosition(t, []string{
		`#!/usr/bin/env node

import { foo } from "./a";

foo`,
	}, nil /*preferences*/)
}
