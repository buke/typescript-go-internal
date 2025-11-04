package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoForJSDocCodefence(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**
 * @example
 * ` + "`" + `` + "`" + `` + "`" + `
 * 1 + 2
 * ` + "`" + `` + "`" + `` + "`" + `
 */
function fo/*1*/o() {
    return '2';
}
/**
 * @example
 * ` + "`" + `` + "`" + `
 * 1 + 2
 * ` + "`" + `
 */
function bo/*2*/o() {
    return '2';
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineHover(t)
}
