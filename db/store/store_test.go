package store

import (
	"testing"
)

func TestNewStore(t *testing.T) {
	memStore := NewStore()
	if memStore.(*MemoryStore).riskStore == nil {
		t.Fatalf("Failed to initialize the store")
	}
}
