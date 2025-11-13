package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_jsx1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @Filename: /node_modules/react/index.d.ts
export const React: any;
// @Filename: /a.tsx
[|<this>|]</this>
// @Filename: /Foo.tsx
export const Foo = 0;
// @Filename: /c.tsx
import { React } from "react";
<Foo />;
// @Filename: /d.tsx
import { Foo } from "./Foo";
<Foo />;`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToFile(t, "/a.tsx")
	f.VerifyImportFixAtPosition(t, []string{}, nil /*preferences*/)
	f.GoToFile(t, "/c.tsx")
	f.VerifyImportFixAtPosition(t, []string{
		`import { React } from "react";
import { Foo } from "./Foo";
<Foo />;`,
	}, nil /*preferences*/)
	f.GoToFile(t, "/d.tsx")
	f.VerifyImportFixAtPosition(t, []string{
		`import { React } from "react";
import { Foo } from "./Foo";
<Foo />;`,
	}, nil /*preferences*/)
}
