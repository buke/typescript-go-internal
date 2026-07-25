package project

import (
	"context"

	"github.com/buke/typescript-go-internal/pkg/diagnostics"
	"github.com/buke/typescript-go-internal/pkg/locale"
	"github.com/buke/typescript-go-internal/pkg/lsp/lsproto"
)

type Client interface {
	WatchFiles(ctx context.Context, id WatcherID, watchers []*lsproto.FileSystemWatcher) error
	UnwatchFiles(ctx context.Context, id WatcherID) error
	RefreshDiagnostics(ctx context.Context) error
	PublishDiagnostics(ctx context.Context, params *lsproto.PublishDiagnosticsParams) error
	RefreshInlayHints(ctx context.Context) error
	RefreshCodeLens(ctx context.Context) error
	ProgressStart(message *diagnostics.Message, args ...any)
	ProgressFinish(message *diagnostics.Message, args ...any)
	SendTelemetry(ctx context.Context, telemetry lsproto.TelemetryEvent) error
	IsActive() bool
	// SetLocale updates the locale used for diagnostic messages.
	SetLocale(locale string)
	// GetLocale returns the current display locale for diagnostic messages.
	GetLocale() locale.Locale
}
