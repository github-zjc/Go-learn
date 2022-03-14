package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var (
	ErrPoolClose = errors.New("pool has been closed")
)

//Pool资源池，让goroutine们来安全的共享资源
type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	mtx       sync.Mutex
	closed    bool
}

func New(factory func() (io.Closer, error), size uint) (*Pool, error) {
	//判断size是否合法
	if size <= 0 {
		return nil, errors.New("invalid size for the resources pool")
	}
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, size),
		closed:    false,
	}, nil

}

//获取资源
func (p *Pool) AcquireResources() (io.Closer, error) {

	select {
	case resources, ok := <-p.resources: //释放使用完的resource资源
		if !ok { //p.resources已经关闭
			return nil, ErrPoolClose
		}
		fmt.Println("get resources from the pool")
		return resources, nil
	default:
		fmt.Println("acquire new resources")
		return p.factory()

	}
}

//操作资源（释放资源或将资源放入resource中）
func (p *Pool) ReleaseResources(resources io.Closer) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		resources.Close()
		return
	}

	select {
	case p.resources <- resources: //如果能够丢尽资源池
		fmt.Println("release resources back to the pool")
	default:
		fmt.Println("release resources closed")
		resources.Close()
	}
}

//关闭资源池
func (p *Pool) Close() {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)
	for resources := range p.resources {
		resources.Close()
	}
}
