package util

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	if config, err := LoadConfig("../"); err == nil {
		if len(config.ServerAddress) < 0 {
			t.Fatal("Config loading failed")
		}
	} else {
		t.Fatalf("Config Loading failed with error %v", err)
	}
}
