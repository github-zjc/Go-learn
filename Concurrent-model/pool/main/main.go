package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"pool"
	_ "strconv"
	"sync"
	"sync/atomic"
	"time"
)

//DBConnection 是我们定义的一个资源
type DBConnection struct {
	id int32
}

func (D DBConnection) Close() error {
	fmt.Println("database closed , #" + fmt.Sprint(D.id))
	return nil
}

var counter int32

func Factory() (io.Closer, error) {

	atomic.AddInt32(&counter, 1)
	return &DBConnection{id: counter}, nil
}

var wg sync.WaitGroup

//执行查询
func performQuery(query int, pool *pool.Pool) {
	defer wg.Done()
	resources, err := pool.AcquireResources()
	if err != nil {
		fmt.Println(err)
	}
	defer pool.ReleaseResources(resources)
	t := rand.Int()%10 + 1 //t [1,10]
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("finish query " + fmt.Sprint(query))
}

func main() {
	pool, err := pool.New(Factory, 5)
	if err != nil {
		log.Fatalln(err)
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)   //如果不加这个休眠，会导致全部创建新的resource
		go performQuery(i, pool)
	}
	wg.Wait()

	pool.Close()
}
