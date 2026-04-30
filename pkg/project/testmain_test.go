package project_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/core"
	"github.com/buke/typescript-go-internal/pkg/testutil/baseline"
)

func TestMain(m *testing.M) {
	core.ApplyDebugStackLimit()
	defer baseline.Track()()
	m.Run()
}
