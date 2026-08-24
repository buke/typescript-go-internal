package api

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/ast"
	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
	"github.com/buke/typescript-go-internal/pkg/parser"
	"github.com/buke/typescript-go-internal/pkg/tspath"
	"gotest.tools/v3/assert"
)

func TestToAPITextEditsUsesOriginalCoordinates(t *testing.T) {
	t.Parallel()

	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: "/app.vue",
		Path:     tspath.Path("/app.vue"),
	}, "const transformed = true;", core.ScriptKindTS)
	sourceFile.SetContentMapperInfo(ast.ContentMapperSourceFileInfo{
		OriginalText:  "😀\nabc",
		ContentMapper: "mapper",
	})

	edits := toAPITextEdits(sourceFile, []*lsproto.TextEdit{{
		Range: lsproto.Range{
			Start: lsproto.Position{Line: 1, Character: 1},
			End:   lsproto.Position{Line: 1, Character: 2},
		},
		NewText: "x",
	}})

	assert.DeepEqual(t, edits, []*TextEdit{{
		Pos:     4,
		End:     5,
		NewText: "x",
	}})
}
