package api

import (
	"testing"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()
	if cache == nil {
		t.Error("Expected new cache to be created")
		return
	}
	if len(cache.data) != 0 {
		t.Error("Expected new cache to be empty")
		return
	}
}

func TestCache_SetAndGet(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := "testValue"

	cache.Set(key, value)
	retrievedValue, ok := cache.Get(key)
	if !ok {
		t.Error("Expected to retrieve value from cache")
	}
	if retrievedValue != value {
		t.Errorf("Expected value %v, got %v", value, retrievedValue)
	}
}

func TestCache_GetNonExistentKey(t *testing.T) {
	cache := NewCache()
	_, ok := cache.Get("nonExistentKey")
	if ok {
		t.Error("Expected to not retrieve value for non-existent key")
	}
}
func TestCache_Delete(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := "testValue"

	cache.Set(key, value)
	cache.Delete(key)
	_, ok := cache.Get(key)
	if ok {
		t.Error("Expected key to be deleted from cache")
	}
}
