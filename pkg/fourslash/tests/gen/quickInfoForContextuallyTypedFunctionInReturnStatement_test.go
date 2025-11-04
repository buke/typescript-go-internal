package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestQuickInfoForContextuallyTypedFunctionInReturnStatement(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Accumulator {
    clear(): void;
    add(x: number): void;
    result(): number;
}

function makeAccumulator(): Accumulator {
    var sum = 0;
    return {
        clear: function () { sum = 0; },
        add: function (val/**/ue) { sum += value; },
        result: function () { return sum; }
    };
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyQuickInfoAt(t, "", "(parameter) value: number", "")
}
