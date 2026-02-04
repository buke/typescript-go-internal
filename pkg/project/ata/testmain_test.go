package ata_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/testutil/baseline"
)

func TestMain(m *testing.M) {
	defer baseline.Track()()
	m.Run()
}
