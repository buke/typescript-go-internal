package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSyntaxErrorAfterImport1(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare module "extmod" {
  namespace IntMod {
    class Customer {
      constructor(name: string);
    }
  }
}
import ext = require('extmod');
import int = ext.IntMod;
var x = new int/*0*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "0")
	f.Insert(t, ".")
}
