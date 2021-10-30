package go_pool

type Worker struct {
	callback func()
	queue    chan Job
}

func NewWorker(queue chan Job, callback func()) *Worker {
	return &Worker{
		callback: callback,
		queue:    queue,
	}
}

func (w *Worker) Start() {
	go func() {
		defer w.callback()

		for {
			job, ok := <-w.queue
			if !ok || job == nil {
				return
			}

			job()
		}
	}()
}

func (w *Worker) Stop() {
	w.queue <- nil
}