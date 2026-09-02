package sourcemap

import "github.com/buke/typescript-go-internal/v7/pkg/core"

type Source interface {
	Text() string
	FileName() string
	ECMALineMap() []core.TextPos
}
