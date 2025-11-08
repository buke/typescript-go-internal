package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestInlayHintsInteractiveTemplateLiteralTypes(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function getTemplateLiteral1(): ` + "`" + `${string},${string}` + "`" + `;
const lit1 = getTemplateLiteral1();
declare function getTemplateLiteral2(): ` + "`" + `\${${string},${string}` + "`" + `;
const lit2 = getTemplateLiteral2();
declare function getTemplateLiteral3(): ` + "`" + `start${string}\${,$${string}end` + "`" + `;
const lit3 = getTemplateLiteral3();
declare function getTemplateLiteral4(): ` + "`" + `${string}\` + "`" + `,${string}` + "`" + `;
const lit4 = getTemplateLiteral4();`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{IncludeInlayVariableTypeHints: true})
}
