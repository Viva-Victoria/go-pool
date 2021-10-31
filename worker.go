package pool

type Worker struct {
	callback func()
	queue    chan Job
	id       int
}

var (
	QuitJob   Job
	idCounter = 0
)

func NewWorker(queue chan Job, callback func()) *Worker {
	idCounter++
	return &Worker{
		callback: callback,
		queue:    queue,
		id:       idCounter,
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			job, ok := <-w.queue
			if !ok || job == nil {
				return
			}

			job(w.id)
			w.callback()
		}
	}()
}
