package pokecache

type Cache struct {
	CacheMap map[string][]byte
}

func (cache Cache) NewCache() {
	cache.CacheMap = make(map[string][]byte)
}

func (cache Cache) Add(key string, val []byte) {
	cache.CacheMap[key] = val
}

func (cache Cache) Get(key string) ([]byte, bool) {
	out, ok := cache.CacheMap[key]
	if ok {
		return out, ok
	}
	return nil, false
}
