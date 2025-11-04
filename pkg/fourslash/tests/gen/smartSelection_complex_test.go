package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSmartSelection_complex(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type X<T, P> = IsExactlyAny<P> extends true ? T : ({ [K in keyof P]: IsExactlyAny<P[K]> extends true ? K extends keyof T ? T[K] : P[/**/K] : P[K]; } & Pick<T, Exclude<keyof T, keyof P>>)`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineSelectionRanges(t)
}
