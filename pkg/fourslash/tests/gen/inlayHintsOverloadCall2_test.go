package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestInlayHintsOverloadCall2(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type HasID = {
    id: number;
}

type Numbers = {
    n: number[];
}

declare function func(bad1: number, bad2: HasID): void;
declare function func(ok_1: Numbers, ok_2: HasID): void;

func(
    { n: [1, 2, 3] },
    {
        id: 1,
    },
);`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{IncludeInlayParameterNameHints: lsutil.IncludeInlayParameterNameHintsAll})
}
