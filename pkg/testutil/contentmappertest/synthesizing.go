package contentmappertest

import (
	"context"
	"fmt"

	"github.com/buke/typescript-go-internal/pkg/contentmapper"
	"github.com/buke/typescript-go-internal/pkg/json"
	"github.com/buke/typescript-go-internal/pkg/spanmap"
)

const synthesizedOutput = "export const el = jsxRuntime(Widget);\n"

type synthesizingHandler struct{ noNotifications }

func (synthesizingHandler) HandleRequest(ctx context.Context, method string, params json.Value) (any, error) {
	switch method {
	case contentmapper.MethodInitialize:
		return initializeResult("mapper"), nil
	case contentmapper.MethodTransform:
		var p contentmapper.TransformParams
		if err := json.Unmarshal(params, &p); err != nil {
			return nil, err
		}
		mappings, err := spanmap.New(nil).Marshal()
		if err != nil {
			return nil, err
		}
		return contentmapper.TransformResult{MappedOutput: contentmapper.MappedOutput{
			Text:      synthesizedOutput,
			Extension: ".ts",
			Mappings:  json.Value(mappings),
		}}, nil
	default:
		return nil, fmt.Errorf("contentmappertest: unexpected method %q", method)
	}
}
