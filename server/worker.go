package server

import "io"

type worker struct {
}

func (w *worker) Init() {

}

func (w *worker) Start() {

}

func (w *worker) Stop() {

}

func (w *worker) Shutdown() {

}

func (w *worker) Output(out io.Writer) {

}

// NewWorker creates a new worker instance
func NewWorker() Server {
	return &worker{}
}
