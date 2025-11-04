package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsdocTypedefTagGoToDefinition(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @allowNonTsExtensions: true
// @Filename: jsdocCompletion_typedef.js
/**
 * @typedef {Object} Person
 * @property {string} /*1*/personName
 * @property {number} personAge
 */

/**
 * @typedef {{ /*2*/animalName: string, animalAge: number }} Animal
 */

/** @type {Person} */
var person; person.[|personName/*3*/|]

/** @type {Animal} */
var animal; animal.[|animalName/*4*/|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, true, "3", "4")
}
