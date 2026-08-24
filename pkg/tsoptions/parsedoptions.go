package tsoptions

import (
	"github.com/buke/typescript-go-internal/pkg/contentmapper"
	"github.com/buke/typescript-go-internal/pkg/core"
)

type ParsedOptions struct {
	CompilerOptions *core.CompilerOptions `json:"compilerOptions"`
	WatchOptions    *core.WatchOptions    `json:"watchOptions"`
	TypeAcquisition *core.TypeAcquisition `json:"typeAcquisition"`

	FileNames         []string                 `json:"fileNames"`
	ProjectReferences []*core.ProjectReference `json:"projectReferences"`
	ContentMappers    []*contentmapper.Mapper  `json:"contentMappers"`
}
