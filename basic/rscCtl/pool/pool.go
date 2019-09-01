package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m        sync.Mutex
	//可以管理任意实现了Closer接口的资源类型
	resource chan io.Closer
	maxSize  int
	usedSize int
	factory  func() (io.Closer, error)
	closed   bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

//创建管理池
func New(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size < 0 {
		return nil, errors.New("Size too small.")
	}

	return &Pool{
		factory:  fn,
		resource: make(chan io.Closer, size),
		maxSize: size,
	}, nil
}

//get resource
func (p *Pool) Acquire() (io.Closer, error) {
	p.m.Lock()
	defer p.m.Unlock()

	select {
	case r, ok := <-p.resource:
		log.Println("Acquire:", "Shard Resource")
		if !ok {
			//管道已经关闭
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		if p.usedSize < p.maxSize {
			p.usedSize++
			log.Printf("Acquire:" + "New Resource." +
				"resource present size/max: %d/%d\n", p.usedSize, p.maxSize)
			return p.factory()
		} else {
			//log.Printf("Acquire:" +
			//	"block for pool's dry, present size: %d/%d", p.usedSize, p.maxSize)
			return nil, nil
		}
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	p.usedSize--

	select {
	//资源放回队列
	case p.resource <- r:
		log.Println("Release:", "into queue")
	//队列满的情况关闭资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	//关闭通道
	close(p.resource)
	//关闭通道资源
	for r := range p.resource{
		p.usedSize--
		r.Close()
	}
}
