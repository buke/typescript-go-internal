package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestInlayHintsTypeMatchesName(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type Client = {};
function getClient(): Client { return {}; };
const client = getClient();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{IncludeInlayVariableTypeHints: true, IncludeInlayVariableTypeHintsWhenTypeMatchesName: false})
}
