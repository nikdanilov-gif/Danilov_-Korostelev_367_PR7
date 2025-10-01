package main

import "time"

type Cache struct {
	data map[string]interface{}
	ttl  map[string]time.Time
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
		ttl:  make(map[string]time.Time),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.data[key] = value
	c.ttl[key] = time.Now().Add(ttl)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	expire, exists := c.ttl[key]
	if !exists {
		return nil, false
	}
	
	if time.Now().After(expire) {
		delete(c.data, key)
		delete(c.ttl, key)
		return nil, false
	}
	
	return c.data[key], true
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
	delete(c.ttl, key)
}

func main() {
	cache := NewCache()
	
	cache.Set("имя", "Сережа", time.Second*2)
	cache.Set("число", 52, time.Second*5)
	cache.Set("город", "Москва", time.Second*10)
	
	if val, ok := cache.Get("имя"); ok {
		println("Найдено имя:", val.(string))
	}
	
	if val, ok := cache.Get("число"); ok {
		println("Найдено число:", val.(int))
	}
	
	time.Sleep(time.Second * 3)
	
	if _, ok := cache.Get("имя"); !ok {
		println("Запись 'имя' удалена (истек срок)")
	}
	
	if val, ok := cache.Get("число"); ok {
		println("Число еще в кэше:", val.(int))
	}
	
	cache.Delete("город")
	println("Запись 'город' удалена вручную")
}