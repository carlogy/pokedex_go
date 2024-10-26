package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Microsecond)

	if cache.cacheEntry == nil {
		t.Error("cache is nil")
	}
}

func TestAddToCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "https://testURL.com",
			inputVal: []byte("Value 1"),
		}, {
			inputKey: "",
			inputVal: []byte("Value 2"),
		},
	}

	for _, entry := range cases {
		cache.Add(entry.inputKey, entry.inputVal)
		actual, ok := cache.Get(entry.inputKey)
		if !ok {
			t.Errorf("%s not found", entry.inputKey)
			continue
		}
		if string(actual) != string(entry.inputVal) {
			t.Errorf("%s does not match %s",
				string(actual),
				string(entry.inputVal),
			)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "Key1"
	cache.Add(key, []byte("Value 1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("%s should have been reaped", key)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "Key1"
	cache.Add(key, []byte("Value 1"))

	time.Sleep(time.Millisecond)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("%s should not have been reaped", key)
	}
}
