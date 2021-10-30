package pool

type Job func()

type Pool interface {
	Add(job Job)
	Size() int
	Wait()
}
