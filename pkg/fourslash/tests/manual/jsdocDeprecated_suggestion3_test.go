package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestJsdocDeprecated_suggestion3(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `interface RequestOptions {
    /** @deprecated use signal instead */
    timeout?: number;
}
declare function request(url: string, opts: RequestOptions): void;

request("/api", { [|timeout|]: 5000 });
declare const opts: RequestOptions;
opts.[|timeout|];`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifySuggestionDiagnostics(t, []*lsproto.Diagnostic{
		{
			Code:    &lsproto.IntegerOrString{Integer: new(int32(6385))},
			Message: lsproto.StringOrMarkupContent{String: new("'timeout' is deprecated.")},
			Tags:    &[]lsproto.DiagnosticTag{lsproto.DiagnosticTagDeprecated},
			Range:   f.Ranges()[0].LSRange,
		},
		{
			Code:    &lsproto.IntegerOrString{Integer: new(int32(6385))},
			Message: lsproto.StringOrMarkupContent{String: new("'timeout' is deprecated.")},
			Tags:    &[]lsproto.DiagnosticTag{lsproto.DiagnosticTagDeprecated},
			Range:   f.Ranges()[1].LSRange,
		},
	})
}
