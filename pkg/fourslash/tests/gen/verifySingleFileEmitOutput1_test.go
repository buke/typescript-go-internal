package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestVerifySingleFileEmitOutput1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: verifySingleFileEmitOutput1_file0.ts
export class A {
}
export class Z {
}
// @Filename: verifySingleFileEmitOutput1_file1.ts
import f = require("./verifySingleFileEmitOutput1_file0");
var /**/b = new f.A();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "var b: f.A", "")
}
