package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	// Логика выталкивания из-за размера очереди.
	t.Run("pushing out capacity", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("aaa", 100) // [aaa]
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200) // [bbb, aaa]
		require.False(t, wasInCache)

		wasInCache = c.Set("ccc", 300) // [ccc, bbb, aaa]
		require.False(t, wasInCache)

		wasInCache = c.Set("ddd", 400) // [ddd, ccc, bbb]
		require.False(t, wasInCache)

		_, ok := c.Get("aaa")
		require.False(t, ok)
	})

	// Логика выталкивания давно используемых элементов.
	t.Run("pushing out time using", func(t *testing.T) {
		c := NewCache(3)

		wasInCache := c.Set("aaa", 100) // [aaa]
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200) // [bbb, aaa]
		require.False(t, wasInCache)

		wasInCache = c.Set("ccc", 300) // [ccc, bbb, aaa]
		require.False(t, wasInCache)

		_, ok := c.Get("aaa") // [aaa, ccc, bbb]
		require.True(t, ok)

		wasInCache = c.Set("bbb", 400) // [bbb, aaa, ccc]
		require.True(t, wasInCache)

		wasInCache = c.Set("ddd", 500) // [ddd, bbb, aaa]
		require.False(t, wasInCache)

		_, ok = c.Get("ccc")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
