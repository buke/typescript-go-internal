package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

// https://github.com/microsoft/typescript-go/issues/3944
func TestCodeFixImportNonTextualSpecifierText(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = "// @Filename: /a.ts\n" +
		"import type { A } from `./${myFolder}/${myFile}`;\n" +
		"\n" +
		"new A/**/()"
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "")
	f.VerifyImportFixAtPosition(t, []string{
		"import { A } from `./${myFolder}/${myFile}`;\n" +
			"\n" +
			"new A()",
	}, nil /*preferences*/)
}
