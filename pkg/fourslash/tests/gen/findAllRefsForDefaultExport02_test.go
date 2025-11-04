package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFindAllRefsForDefaultExport02(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/export default function /*2*/DefaultExportedFunction() {
    return /*3*/DefaultExportedFunction;
}

var x: typeof /*4*/DefaultExportedFunction;

var y = /*5*/DefaultExportedFunction();

/*6*/namespace /*7*/DefaultExportedFunction {
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3", "4", "5", "6", "7")
}
