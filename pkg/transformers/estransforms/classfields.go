package estransforms

import (
	"github.com/buke/typescript-go-internal/pkg/ast"
	"github.com/buke/typescript-go-internal/pkg/transformers"
)

type classFieldsTransformer struct {
	transformers.Transformer
}

func (ch *classFieldsTransformer) visit(node *ast.Node) *ast.Node {
	return node // !!!
}

func newClassFieldsTransformer(opts *transformers.TransformOptions) *transformers.Transformer {
	tx := &classFieldsTransformer{}
	return tx.NewTransformer(tx.visit, opts.Context)
}
