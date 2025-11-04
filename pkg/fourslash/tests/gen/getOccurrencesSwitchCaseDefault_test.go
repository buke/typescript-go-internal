package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesSwitchCaseDefault(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|switch|] (10) {
    [|case|] 1:
    [|case|] 2:
    [|case|] 4:
    [|case|] 8:
        foo: switch (20) {
            case 1:
            case 2:
                break;
            default:
                break foo;
        }
    [|case|] 0xBEEF:
    [|default|]:
        [|break|];
    [|case|] 16:
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
