package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestTsxIncrementalServer(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: es5
/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.MarkTestAsStradaServer()
	f.GoToMarker(t, "")
	f.Insert(t, "<")
	f.Insert(t, "div")
	f.Insert(t, " ")
	f.Insert(t, " id")
	f.Insert(t, "=")
	f.Insert(t, "\"foo")
	f.Insert(t, "\"")
	f.Insert(t, ">")
}
