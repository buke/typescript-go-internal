package core

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=ScriptKind -output=scriptkind_stringer_generated.go
//go:generate go run mvdan.cc/gofumpt@latest -w scriptkind_stringer_generated.go

type ScriptKind int32

const (
	ScriptKindUnknown ScriptKind = iota
	ScriptKindJS
	ScriptKindJSX
	ScriptKindTS
	ScriptKindTSX
	ScriptKindExternal
	ScriptKindJSON
	/**
	 * Used on extensions that doesn't define the ScriptKind but the content defines it.
	 * Deferred extensions are going to be included in all project contexts.
	 */
	ScriptKindDeferred
)
