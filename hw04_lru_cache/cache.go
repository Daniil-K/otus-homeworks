package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set Добавить значение в кэш по ключу
func (cache *lruCache) Set(key Key, value interface{}) bool {
	item, ok := cache.items[key]

	if ok {
		item.Value = value
		cache.queue.MoveToFront(item)
		return true
	}

	listItem := cache.queue.PushFront(value)
	cache.items[key] = listItem

	if cache.queue.Len() > cache.capacity {
		lastItem := cache.queue.Back()
		cache.queue.Remove(lastItem)

		for k, val := range cache.items {
			if val == lastItem {
				delete(cache.items, k)
			}
		}
	}

	return false
}

// Get Получить значение из кэша по ключу
func (cache *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := cache.items[key]

	if ok {
		cache.queue.MoveToFront(item)
		return item.Value, true
	}

	return nil, false
}

// Clear Очистить кэш
func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
