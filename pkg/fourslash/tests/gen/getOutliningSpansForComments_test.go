package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOutliningSpansForComments(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: es5
[|/*
    Block comment at the beginning of the file before module:
        line one of the comment
        line two of the comment
        line three
        line four
        line five
*/|]
declare module "m";
[|// Single line comments at the start of the file
// line 2
// line 3
// line 4|]
declare module "n";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.MarkTestAsStradaServer()
	f.VerifyOutliningSpans(t, lsproto.FoldingRangeKindComment)
}
