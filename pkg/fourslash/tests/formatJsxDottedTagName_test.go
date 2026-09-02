package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestFormatJsxDottedTagName(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
const x = (
<a-b.c>
<a-b.c></a-b.c>
</a-b.c>
);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, `const x = (
    <a-b.c>
        <a-b.c></a-b.c>
    </a-b.c>
);`)
}
