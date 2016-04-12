package cache

import "testing"

func TestCache(t *testing.T) {
	InitCache("./cache_config.conf")
	Set("abc", "def")
	val, err := GetString("abc")
	t.Logf("val = %s, err = %v", val, err)
}