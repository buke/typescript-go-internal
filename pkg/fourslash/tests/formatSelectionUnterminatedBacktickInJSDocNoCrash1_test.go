package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestFormatSelectionUnterminatedBacktickInJSDocNoCrash1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = "// @allowJs: true\n" +
		"// @filename: index.js\n" +
		"export class Manifest {\n" +
		"  /**\n" +
		"   * @template {ExtensionType} ExtType\n" +
		"   * @param {ExtType} extType - `dri\n" +
		"/*begin*/   * @param {string} extName\n" +
		"   */\n" +
		"  setExtension(extType, extName, extData) {\n" +
		"        const data = _.cloneDeep(extData);\n" +
		"    this.#data[`${extType}s`][extName] = data;\n" +
		"    return data;\n" +
		"  }\n" +
		"/*end*/\n" +
		"}\n" +
		"\n"
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatSelection(t, "begin", "end")
	f.VerifyCurrentFileContent(t, "export class Manifest {\n"+
		"  /**\n"+
		"   * @template {ExtensionType} ExtType\n"+
		"   * @param {ExtType} extType - `dri\n"+
		"   * @param {string} extName\n"+
		"   */\n"+
		"    setExtension(extType, extName, extData) {\n"+
		"        const data = _.cloneDeep(extData);\n"+
		"        this.#data[`${extType}s`][extName] = data;\n"+
		"        return data;\n"+
		"    }\n"+
		"\n"+
		"}\n"+
		"\n")
}
