package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOutliningSpansDepthElseIf(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else if (1)[| {
    1;
}|] else[| {
    1;
}|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyOutliningSpans(t)
}
