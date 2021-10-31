package pool

type Job func(workerId int)

type Pool interface {
	Add(job Job)
	Size() int
	Wait()
}
