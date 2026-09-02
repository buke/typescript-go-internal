package emittestutil

import (
	"strings"
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/ast"
	"github.com/buke/typescript-go-internal/v7/pkg/core"
	"github.com/buke/typescript-go-internal/v7/pkg/printer"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil/parsetestutil"
	"gotest.tools/v3/assert"
)

// Checks that pretty-printing the given file matches the expected output.
func CheckEmit(t *testing.T, emitContext *printer.EmitContext, file *ast.SourceFile, expected string) {
	t.Helper()
	printer := printer.NewPrinter(
		printer.PrinterOptions{
			NewLine: core.NewLineKindLF,
		},
		printer.PrintHandlers{},
		emitContext,
	)
	text := printer.EmitSourceFile(file)
	actual := strings.TrimSuffix(text, "\n")
	assert.Equal(t, expected, actual)
	file2 := parsetestutil.ParseTypeScript(text, file.LanguageVariant == core.LanguageVariantJSX)
	parsetestutil.CheckDiagnosticsMessage(t, file2, "error on reparse: ")
}
