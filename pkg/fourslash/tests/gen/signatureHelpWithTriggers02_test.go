package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpWithTriggers02(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function foo<T>(x: T, y: T): T;
declare function bar<U>(x: U, y: U): U;

foo(bar/*1*/)`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.Insert(t, "(")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "bar(x: unknown, y: unknown): unknown"})
	f.Backspace(t, 1)
	f.Insert(t, "<")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "bar<U>(x: U, y: U): U"})
	f.Backspace(t, 1)
	f.Insert(t, ",")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "foo(x: <U>(x: U, y: U) => U, y: <U>(x: U, y: U) => U): <U>(x: U, y: U) => U"})
	f.Backspace(t, 1)
}
