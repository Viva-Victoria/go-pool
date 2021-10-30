package pool

import (
	"sync"
)

type FixedPool struct {
	queue     chan Job
	workers   []Worker
	waitGroup sync.WaitGroup
}

func NewFixedPool(size int) (*FixedPool, error) {
	if size < 1 || size > 65535 {
		return nil, ErrIncorrectSize
	}

	pool := &FixedPool{
		queue:   make(chan Job),
		workers: make([]Worker, size),
	}
	for i := 0; i < size; i++ {
		pool.waitGroup.Add(1)
		pool.workers[i] = *NewWorker(pool.queue, func() {
			pool.waitGroup.Done()
		})
	}

	return pool, nil
}

func (p *FixedPool) Size() int {
	return len(p.workers)
}

func (p *FixedPool) Add(job Job) {
	p.waitGroup.Add(1)
	p.queue <- job
}

func (p *FixedPool) Wait() {
	p.waitGroup.Wait()
	close(p.queue)
}
