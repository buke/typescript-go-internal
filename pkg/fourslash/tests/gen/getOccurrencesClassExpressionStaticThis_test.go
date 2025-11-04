package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	. "github.com/buke/typescript-go-internal/pkg/fourslash/tests/util"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOccurrencesClassExpressionStaticThis(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `var x = class C {
    public x;
    public y;
    public z;
    public staticX;
    constructor() {
        this;
        this.x;
        this.y;
        this.z;
    }
    foo() {
        this;
        () => this;
        () => {
            if (this) {
                this;
            }
        }
        function inside() {
            this;
            (function (_) {
                this;
            })(this);
        }
        return this.x;
    }

    static bar() {
        [|this|];
        [|this|].staticX;
        () => [|this|];
        () => {
            if ([|this|]) {
                [|this|];
            }
        }
        function inside() {
            this;
            (function (_) {
                this;
            })(this);
        }
    }
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}
