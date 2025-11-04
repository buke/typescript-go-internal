package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSmartSelection_emptyRanges(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `class HomePage {
  componentDidMount(/*1*/) {
    if (this.props.username/*2*/) {
      return '/*3*/';
    }
  }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineSelectionRanges(t)
}
