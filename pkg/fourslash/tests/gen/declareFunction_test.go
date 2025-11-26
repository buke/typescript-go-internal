package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestDeclareFunction(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: index.ts
declare function`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "",
			Preferences: nil,
			Exact:       PtrTo([]*lsproto.SymbolInformation{}),
		},
	})
}
