package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_importType1(t *testing.T) {
	t.Parallel()
	t.Skip()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @verbatimModuleSyntax: true
// @module: es2015
// @Filename: /exports.ts
export default someValue = 0;
export function Component() {}
export interface ComponentProps {}
// @Filename: /a.ts
import { Component } from "./exports.js";
interface MoreProps extends /*a*/ComponentProps {}
// @Filename: /b.ts
import someValue from "./exports.js";
interface MoreProps extends /*b*/ComponentProps {}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "a")
	f.VerifyImportFixAtPosition(t, []string{
		`import { Component, type ComponentProps } from "./exports.js";
interface MoreProps extends ComponentProps {}`,
	}, nil /*preferences*/)
	f.GoToMarker(t, "b")
	f.VerifyImportFixAtPosition(t, []string{
		`import someValue, { type ComponentProps } from "./exports.js";
interface MoreProps extends ComponentProps {}`,
	}, nil /*preferences*/)
}
