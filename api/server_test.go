package api

import (
	pstore "github.com/vivekgeorgemathew/aw/db/store"
	"testing"
)

func TestNewServer(t *testing.T) {
	store := pstore.NewStore()
	server := NewServer(store)
	routesInfo := server.router.Routes()
	if len(routesInfo) != 3 {
		t.Fatalf("Expected 3 routes but got %d", len(routesInfo))
	}
}
