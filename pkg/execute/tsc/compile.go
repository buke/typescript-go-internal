package tsc

import (
	"io"
	"time"

	"github.com/buke/typescript-go-internal/pkg/ast"
	"github.com/buke/typescript-go-internal/pkg/collections"
	"github.com/buke/typescript-go-internal/pkg/compiler"
	"github.com/buke/typescript-go-internal/pkg/execute/incremental"
	"github.com/buke/typescript-go-internal/pkg/tspath"
	"github.com/buke/typescript-go-internal/pkg/vfs"
)

type System interface {
	Writer() io.Writer
	FS() vfs.FS
	DefaultLibraryPath() string
	GetCurrentDirectory() string
	WriteOutputIsTTY() bool
	GetWidthOfTerminal() int
	GetEnvironmentVariable(name string) string

	Now() time.Time
	SinceStart() time.Duration
}

type ExitStatus int

const (
	ExitStatusSuccess                              ExitStatus = 0
	ExitStatusDiagnosticsPresent_OutputsGenerated  ExitStatus = 1
	ExitStatusDiagnosticsPresent_OutputsSkipped    ExitStatus = 2
	ExitStatusInvalidProject_OutputsSkipped        ExitStatus = 3
	ExitStatusProjectReferenceCycle_OutputsSkipped ExitStatus = 4
	ExitStatusNotImplemented                       ExitStatus = 5
)

type Watcher interface {
	DoCycle()
}

type CommandLineResult struct {
	Status  ExitStatus
	Watcher Watcher
}

type CommandLineTesting interface {
	// Ensure that all emitted files are timestamped in order to ensure they are deterministic for test baseline
	OnEmittedFiles(result *compiler.EmitResult, mTimesCache *collections.SyncMap[tspath.Path, time.Time])
	OnListFilesStart(w io.Writer)
	OnListFilesEnd(w io.Writer)
	OnStatisticsStart(w io.Writer)
	OnStatisticsEnd(w io.Writer)
	OnBuildStatusReportStart(w io.Writer)
	OnBuildStatusReportEnd(w io.Writer)
	OnWatchStatusReportStart()
	OnWatchStatusReportEnd()
	GetTrace(w io.Writer) func(msg string)
	OnProgram(program *incremental.Program)
}

type CompileTimes struct {
	ConfigTime         time.Duration
	ParseTime          time.Duration
	bindTime           time.Duration
	checkTime          time.Duration
	totalTime          time.Duration
	emitTime           time.Duration
	BuildInfoReadTime  time.Duration
	ChangesComputeTime time.Duration
}
type CompileAndEmitResult struct {
	Diagnostics []*ast.Diagnostic
	EmitResult  *compiler.EmitResult
	Status      ExitStatus
	times       *CompileTimes
}
