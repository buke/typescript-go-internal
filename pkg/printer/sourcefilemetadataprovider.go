package printer

import (
	"github.com/buke/typescript-go-internal/pkg/ast"
	"github.com/buke/typescript-go-internal/pkg/tspath"
)

type SourceFileMetaDataProvider interface {
	GetSourceFileMetaData(path tspath.Path) *ast.SourceFileMetaData
}
