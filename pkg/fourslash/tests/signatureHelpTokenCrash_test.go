package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpTokenCrash(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
function foo(a: any, b: any) {

}

foo((/*1*/

/** This is a JSDoc comment */
foo/** More comments*/((/*2*/
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifySignatureHelpWithCases(t, &fourslash.SignatureHelpCase{
		MarkerInput: "1",
		Expected:    nil,
		Context: &lsproto.SignatureHelpContext{
			IsRetrigger:      false,
			TriggerCharacter: PtrTo("("),
			TriggerKind:      lsproto.SignatureHelpTriggerKindTriggerCharacter,
		},
	})
	f.VerifySignatureHelpWithCases(t, &fourslash.SignatureHelpCase{
		MarkerInput: "2",
		Expected:    nil,
		Context: &lsproto.SignatureHelpContext{
			IsRetrigger:      false,
			TriggerCharacter: PtrTo("("),
			TriggerKind:      lsproto.SignatureHelpTriggerKindTriggerCharacter,
		},
	})
}
