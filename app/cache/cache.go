package cache

type CacheStore interface {
	Store(name string)
}

type Cache interface {
	Set(key string, val interface{})
	Get(key string) interface{}
}
