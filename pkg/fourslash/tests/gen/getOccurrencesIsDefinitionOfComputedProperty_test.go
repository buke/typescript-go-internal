package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesIsDefinitionOfComputedProperty(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `let o = { /*1*/["/*2*/foo"]: 12 };
let y = o./*3*/foo;
let z = o['/*4*/foo'];`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3", "4")
}
