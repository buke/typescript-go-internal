package incremental

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/ast"
	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/diagnostics"
	"github.com/buke/typescript-go-internal/pkg/locale"
	"github.com/buke/typescript-go-internal/pkg/parser"
	"gotest.tools/v3/assert"
)

func TestExternalDiagnosticBuildInfoRoundTrip(t *testing.T) {
	t.Parallel()
	file := parser.ParseSourceFile(ast.SourceFileParseOptions{FileName: "/app.vue", Path: "/app.vue"}, "", core.ScriptKindTS)
	diagnostic := ast.NewExternalDiagnostic(file, core.NewTextRange(1, 2), "vue", diagnostics.CategoryWarning, 1001, "mapper warning")

	serialized := astDiagToBuildInfoDiag(diagnostic)
	assert.Equal(t, serialized.source, "vue")
	assert.Equal(t, serialized.messageText, "mapper warning")

	restored := serialized.toDiagnostic(nil, file)
	assert.Equal(t, restored.Source(), "vue")
	assert.Equal(t, restored.Localize(locale.Default), "mapper warning")
}
