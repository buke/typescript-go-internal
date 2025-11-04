package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGoToDefinitionCSSPatternAmbientModule(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @esModuleInterop: true
// @Filename: index.css
/*2a*/html { font-size: 16px; }
// @Filename: types.ts
declare module /*2b*/"*.css" {
  const styles: any;
  export = styles;
}
// @Filename: index.ts
import styles from [|/*1*/"./index.css"|];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "1")
}
