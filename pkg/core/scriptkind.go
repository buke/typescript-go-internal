package core

//go:generate go run golang.org/x/tools/cmd/stringer@latest -type=ScriptKind -output=scriptkind_stringer_generated.go
//go:generate npx dprint fmt scriptkind_stringer_generated.go

type ScriptKind int32

const (
	ScriptKindUnknown ScriptKind = 0
	ScriptKindJS      ScriptKind = 1
	ScriptKindJSX     ScriptKind = 2
	ScriptKindTS      ScriptKind = 3
	ScriptKindTSX     ScriptKind = 4

	// Value 5 is reserved (formerly ScriptKindExternal).

	ScriptKindJSON ScriptKind = 6

	// Value 7 is reserved (formerly ScriptKindDeferred).
)
