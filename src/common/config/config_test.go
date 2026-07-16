package common

import (
	"testing"
)

func TestNewConfigUsesDefaultsWhenEnvFileIsMissing(t *testing.T) {
	cfg := NewConfig()

	if cfg == nil {
		t.Fatal("expected config instance")
	}

	if cfg.App == nil {
		t.Fatal("expected app config")
	}

	if cfg.App.Port == "" {
		t.Fatal("expected default port when env is missing")
	}
}
