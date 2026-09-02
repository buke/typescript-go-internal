package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/v7/pkg/fourslash"
	"github.com/buke/typescript-go-internal/v7/pkg/testutil"
)

func TestQuickinfoVerbosityIncreaseDecrease(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `export const JOB_STATES = ["created", "active", "completed", "failed", "retry", "cancelled", "archive"] as const
export type JobState = (typeof JOB_STATES)[number]
type Color = "default" | "primary" | "secondary" | "success" | "warning" | "danger"
const JobsStateToColor/*a*/: Record<
  JobState,
  {
    color: Color
    label: string
    labelPlural: string
  }
> = {
  created: {
    color: "success",
    label: "Направљен",
    labelPlural: "Направљени",
  },
  active: {
    color: "success",
    label: "Активан",
    labelPlural: "Активни",
  },
  completed: {
    color: "success",
    label: "Успешан",
    labelPlural: "Успешни",
  },
  cancelled: {
    color: "default",
    label: "Отаказан",
    labelPlural: "Отаказни",
  },
  failed: {
    color: "danger",
    label: "Пао",
    labelPlural: "Пали",
  },
  archive: {
    color: "default",
    label: "Архивиран",
    labelPlural: "Архивирани",
  },
  retry: {
    color: "warning",
    label: "Понавља се",
    labelPlural: "Понављају се",
  },
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHoverWithVerbosity(t, map[string][]int{"a": {0, 1, 0}})
}
