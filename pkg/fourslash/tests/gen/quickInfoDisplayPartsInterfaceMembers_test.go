package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoDisplayPartsInterfaceMembers(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I {
    /*1*/property: string;
    /*2*/method(): string;
    (): string;
    new (): I;
}
var iInstance: I;
/*3*/iInstance./*4*/property = /*5*/iInstance./*6*/method();
/*7*/iInstance();
var /*8*/anotherInstance = new /*9*/iInstance();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
