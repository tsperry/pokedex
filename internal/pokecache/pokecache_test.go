package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("first", []byte("original"))
	cache.Add("second", []byte("other"))
	cache.Add("first", []byte("updated"))

	for _, tc := range []struct {
		key   string
		want  string
		found bool
	}{
		{key: "first", want: "updated", found: true},
		{key: "second", want: "other", found: true},
		{key: "missing", found: false},
	} {
		t.Run(tc.key, func(t *testing.T) {
			got, ok := cache.Get(tc.key)
			if ok != tc.found || string(got) != tc.want {
				t.Fatalf("Get(%q) = %q, %v; want %q, %v", tc.key, got, ok, tc.want, tc.found)
			}
			if !ok && got != nil {
				t.Fatal("missing key should return nil")
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	basetime := 5 * time.Second
	cache := NewCache(basetime)
	cache.Add("first", []byte("original"))
	got, ok := cache.Get("first")
	if !ok {
		t.Fatal("unable to get first")
	}
	if string(got) != "original" {
		t.Fatal("bad get")
	}

	time.Sleep(basetime + 5*time.Millisecond)
	got, ok = cache.Get("first")
	if ok {
		t.Fatal("should not have gotten key after reap")
	}
}
