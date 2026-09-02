package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/ls/lsutil"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestAutoImportSymlinkCaseSensitive(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /tsconfig.json
{ "compilerOptions": { "module": "commonjs" } }
// @Filename: /node_modules/.pnpm/mobx@6.0.4/node_modules/MobX/Foo.d.ts
export declare function autorun(): void;
// @Filename: /index.ts
autorun/**/
// @Filename: /utils.ts
import "MobX/Foo";
// @link: /node_modules/.pnpm/mobx@6.0.4/node_modules/MobX -> /node_modules/MobX
// @link: /node_modules/.pnpm/mobx@6.0.4/node_modules/MobX -> /node_modules/.pnpm/cool-mobx-dependent@1.2.3/node_modules/MobX`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.Configure(t, lsutil.UserPreferences{AutoImportEntrypointDirectorySearch: core.TSTrue})
	f.GoToMarker(t, "")
	f.VerifyImportFixAtPosition(t, []string{
		`import { autorun } from "MobX/Foo";

autorun`,
	}, nil /*preferences*/)
}
