package cache

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/coocood/freecache"
)

func GCPause() time.Duration {
	runtime.GC()
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	return stats.Pause[0]
}

func TestFreeCache() {

	n := 3000 * 1000
	debug.SetGCPercent(10)

	freeCache := freecache.NewCache(512 * 1024 * 1024)
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("key%v", i)
		val := make([]byte, 10)
		freeCache.Set([]byte(key), val, 0)
	}
	fmt.Println("GC pause with free cache:", GCPause())
	freeCache = nil

}
