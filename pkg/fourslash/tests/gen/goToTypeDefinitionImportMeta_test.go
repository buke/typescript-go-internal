package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToTypeDefinitionImportMeta(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @module: esnext
// @Filename: foo.ts
/// <reference no-default-lib="true"/>
/// <reference path='./bar.d.ts' />
import.me/*reference*/ta;
//@Filename: bar.d.ts
interface /*definition*/ImportMeta {
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToTypeDefinition(t, "reference")
}
