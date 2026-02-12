package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestOrganizeImportsAttributes3(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `import { A } from "./a";
import { C } from "./a" with {      type: "a" };
import { Z } from "./z";
import { A as D } from "./a" with    { type: "b" };
import { E } from "./a" with { type: /* comment*/ "a"              };
import { F } from "./a" with     {type: "a" };
import { Y } from "./a"   with{ type: "b" /* comment*/};
import { B } from "./a";

export type G = A | B | C | D | E | F | Y | Z;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyOrganizeImports(t,
		`import { A, B } from "./a";
import { C, E, F } from "./a" with { type: "a" };
import { A as D, Y } from "./a" with { type: "b" };
import { Z } from "./z";

export type G = A | B | C | D | E | F | Y | Z;`,
		lsproto.CodeActionKindSourceOrganizeImports,
		nil,
	)
}
