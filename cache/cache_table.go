package cache

// CacheTable represent a table like table of MySQL
type CacheTable struct {
	items map[string]interface{}
}

func (ct *CacheTable) Get(key string) (res interface{}, exist bool) {
	res, exist = ct.items[key]
	return
}
