package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestIncrementalParsingWithJsDoc(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|import a from 'a/aaaaaaa/aaaaaaa/aaaaaa/aaaaaaa';
/**/import b from 'b';
import c from 'c';|]
[|/** @internal */|]
export class LanguageIdentifier[| { }|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOutliningSpans(t)
	f.GoToMarker(t, "")
	f.Backspace(t, 1)
	f.VerifyOutliningSpans(t)
}
