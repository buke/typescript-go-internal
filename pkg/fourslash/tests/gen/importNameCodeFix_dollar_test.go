package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestImportNameCodeFix_dollar(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: esnext
// @moduleResolution: bundler
// @Filename: /node_modules/qwik/index.d.ts
export declare const $: any;
// @Filename: /index.ts
import {} from "qwik";
$/**/`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "")
	f.VerifyImportFixAtPosition(t, []string{
		`import { $ } from "qwik";
$`,
	}, nil /*preferences*/)
}
