package cache

/**
模仿cache2go，自己写一个，主要是练习

写的过程中，不可以看cache2go，只有自己写完或者写不下去了，才可以看，否则就变成了抄代码，没有任何学习的效果
*/

type Cache struct {
}

func New() *Cache {
	return &Cache{}
}

func (cache *Cache) Create(table string) *CacheTable {
	return &CacheTable{}
}
