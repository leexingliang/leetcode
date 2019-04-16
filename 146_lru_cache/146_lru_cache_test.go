package lru_cache

import (
	"fmt"
	"testing"
)

func Test_LRU(t *testing.T) {
	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // returns 1
	lru.Put(3, 3)           // evicts key 2
	fmt.Println(lru.Get(2)) // returns -1 (not found)
	lru.Put(4, 4)           // evicts key 1
	fmt.Println(lru.Get(1)) // returns -1 (not found)
	fmt.Println(lru.Get(3)) // returns 3
	fmt.Println(lru.Get(4)) // returns 4
}

func Test_LRU_PutGet(t *testing.T) {
	lru := Constructor(1)

	lru.Put(1, 2)
	fmt.Println(lru.Get(1)) // returns 2
}

func Test_LRU_PutGetPutGetGet(t *testing.T) {
	lru := Constructor(1)

	lru.Put(2, 1)
	fmt.Println(lru.Get(2)) // returns 1
	lru.Put(3, 2)
	fmt.Println(lru.Get(2)) // returns -1
	fmt.Println(lru.Get(3)) // returns 2
}

func Test_LRU_PPGPPG(t *testing.T) {
	lru := Constructor(2)

	lru.Put(2, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(2)) // returns 1
	lru.Put(1, 4)
	lru.Put(4, 1)
	fmt.Println(lru.Get(2)) // returns -1
	fmt.Println(lru.Get(3)) // returns -1
}
