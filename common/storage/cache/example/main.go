package main

import (
	"log"

	"github.com/pjh130/go/common/storage/cache"
)

func main() {
	log.Println("====================")
	// cache.TestFreeCache()
	// cache.TestGoCache()
	// cache.TestGroupcache()
	cache.TestCache2go()
}
