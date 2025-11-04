package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpTokenCrash2(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
function foo<T, U>(x: string, y: T, z: U) {

}

foo<number,number>/*1*/("hello", 123,456)
`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifySignatureHelp(t, &fourslash.SignatureHelpCase{
		MarkerInput: "1",
		Expected:    nil,
		Context: &lsproto.SignatureHelpContext{
			IsRetrigger:      false,
			TriggerCharacter: PtrTo("("),
			TriggerKind:      lsproto.SignatureHelpTriggerKindTriggerCharacter,
		},
	})
}
