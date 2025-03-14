package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache           *sync.Map
	cacheExpiration time.Duration
}

func New(exp time.Duration) *Cache {
	return &Cache{
		cache:           &sync.Map{},
		cacheExpiration: exp,
	}
}

type cacheEntry struct {
	value     map[string]float32
	expiresAt time.Time
}

func (c *Cache) GetExchangeRates() (map[string]float32, bool) {
	value, ok := c.cache.Load("exchange_rates")
	if !ok {
		return nil, false
	}

	cacheEntry, ok := value.(cacheEntry)
	if !ok || time.Now().After(cacheEntry.expiresAt) {
		return nil, false
	}

	return cacheEntry.value, true
}

func (c *Cache) StoreExchangeRates(rates map[string]float32) {
	c.cache.Store("exchange_rates", cacheEntry{
		value:     rates,
		expiresAt: time.Now().Add(c.cacheExpiration),
	})
}

func (c *Cache) Cleanup() {
	for {
		time.Sleep(c.cacheExpiration / 2)
		c.cache.Range(func(key, value any) bool {
			cacheEntry, ok := value.(cacheEntry)
			if !ok {
				return true
			}

			if time.Now().After(cacheEntry.expiresAt) {
				c.cache.Delete(key)
			}
			return true
		})
	}
}
