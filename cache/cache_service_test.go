package cache_test

import (
	"org.geekible.servicehub/cache"
	"org.geekible.servicehub/configuration"
	"testing"
)

func TestSetKey(t *testing.T) {
	config := configuration.NewServiceConfiguration()
	cacheService := cache.NewCacheService(config)
	if err := cacheService.SetKey("test", "123"); err != nil {
		t.Errorf("failed to insert key into cache: %q", err)
	}
}

func TestGetKey(t *testing.T) {
	config := configuration.NewServiceConfiguration()
	cacheService := cache.NewCacheService(config)
	got, err := cacheService.GetKey("test")
	expected := "123"
	if err != nil {
		t.Errorf("failed to get key 'test' from cache with error: %q", err)
	}

	if got != expected {
		t.Errorf("got does not match expected, got=%q, expected=%q", got, expected)
	}
}

func TestRemoveKey(t *testing.T) {
	config := configuration.NewServiceConfiguration()
	cacheService := cache.NewCacheService(config)
	if err := cacheService.RemoveKey("test"); err != nil {
		t.Errorf("failed to delete ")
	}
}
