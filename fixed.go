package pool

import (
	"sync"
)

type FixedPool struct {
	queue     chan Job
	size      int
	waitGroup sync.WaitGroup
}

const (
	MaxPoolSize = 65536
)

func NewFixedPool(size int) (*FixedPool, error) {
	if size < 1 || size > MaxPoolSize {
		return nil, ErrIncorrectSize
	}

	pool := &FixedPool{
		queue: make(chan Job),
		size:  0,
	}

	_, _ = pool.Expand(size)

	return pool, nil
}

func (p *FixedPool) Collapse(count int) (int, error) {
	if p.size-count < 1 {
		return p.size, ErrIncorrectSize
	}

	for i := 0; i < count; i++ {
		p.queue <- QuitJob
	}

	p.size -= count

	return p.size, nil
}

func (p *FixedPool) Expand(count int) (int, error) {
	if p.size+count > MaxPoolSize {
		return p.size, ErrIncorrectSize
	}

	for i := 0; i < count; i++ {
		worker := NewWorker(p.queue, func() {
			p.waitGroup.Done()
		})
		worker.Start()
	}
	p.size += count

	return p.size, nil
}

func (p *FixedPool) Size() int {
	return p.size
}

func (p *FixedPool) Add(job Job) {
	p.waitGroup.Add(1)
	p.queue <- job
}

func (p *FixedPool) Wait() {
	p.waitGroup.Wait()
	close(p.queue)
}
