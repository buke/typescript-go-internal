package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestSignatureHelpInAdjacentBlockBody(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function foo(...args);

foo(() => {/*1*/}/*2*/)`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.GoToMarker(t, "1")
	f.VerifySignatureHelpPresent(t, &lsproto.SignatureHelpContext{TriggerKind: lsproto.SignatureHelpTriggerKindInvoked})
	f.GoToMarker(t, "2")
	f.VerifySignatureHelpPresent(t, &lsproto.SignatureHelpContext{TriggerKind: lsproto.SignatureHelpTriggerKindInvoked})
}
