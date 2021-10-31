# go-pool

## Install
`go get github.com/VivaVictoria/go-pool`

## Fixed pool
You can create pool with fixed goroutines count. Pool size should be between 1 (minimal size) and 65536 (maximum size).
```go
package main

import (
	"fmt"
	"github.com/Viva-Victoria/go-pool"
	"log"
)

func CreateJob(i int) pool.Job {
	return func(workerId int) {
		log.Println(fmt.Sprintf("job #%d on worker %d", i, workerId))
	}
}

func main() {
	p, err := pool.NewFixedPool(3) //create pool with 3 goroutines
	if err != nil {
		panic(err)
	}

	log.Println(p.Size()) //3

	for i := 0; i < 9; i++ {
		p.Add(CreateJob(i))
	}

	newSize, err := p.Expand(20) //increase pool by 20 goroutines
	if err != nil {
		panic(err)
	}

	log.Println(newSize) //actual is 23

	for i := 0; i < 50; i++ {
		p.Add(CreateJob(i))
	}

	newSize, err = p.Collapse(22) //decrease pool by 22 goroutines
	if err != nil {
		panic(err)
	}

	log.Println(newSize) //actual is 1

	for i := 0; i < 10; i++ {
		p.Add(CreateJob(i))
	}

	p.Wait()
}
```