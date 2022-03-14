package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomicIncCounter(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000000; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func mutexIncCounter(counter *int64, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 10000000; i++ {
		mtx.Lock()
		*counter++
		mtx.Unlock()
	}
}

func ConcurrentAtomic() int64 {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(2)
	go atomicIncCounter(&counter, &wg)
	go atomicIncCounter(&counter, &wg)
	wg.Wait()
	return counter
}

func ConcurrentMutex() int64 {
	var counter int64
	var wg sync.WaitGroup
	var mtx sync.Mutex
	wg.Add(2)
	go mutexIncCounter(&counter, &wg, &mtx)
	go mutexIncCounter(&counter, &wg, &mtx)
	wg.Wait()
	return counter
}

func main() {
	fmt.Println(ConcurrentAtomic())
	fmt.Println(ConcurrentMutex())
}
