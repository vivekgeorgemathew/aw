package store

import "testing"

func TestSave(t *testing.T) {
	runStoreTests(t)
}

func TestGet(t *testing.T) {
	runStoreTests(t)
}

func TestGetAll(t *testing.T) {
	runStoreTests(t)
}

func runStoreTests(t *testing.T) {
	store := NewStore()
	key := "key"
	value := "value"
	_, err := store.Get("invalid")
	if err == nil {
		t.Errorf("Get failed to return proper error for invalid key")
	}
	_, err = store.GetAll()
	if err == nil {
		t.Errorf("GetAll failed to return proper error for empty store")
	}
	if err := store.Save(key, value); err != nil {
		t.Fatalf("Failed to save the value with error %v", err)
	}
	retValue, err := store.Get(key)
	if err != nil {
		t.Fatalf("Failed to save/retrieve the value with error %v", err)
	}
	if retValue != value {
		t.Fatalf("Failed to save the value")
	}
	values, err := store.GetAll()
	if len(values) != 1 {
		t.Fatalf("Expected 1 value but got %d", len(value))
	}
}
