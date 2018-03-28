package main

import "sync"

func main() {
	urls := []string{}
	t := NewTaskRunner(10)  //控制并发度
	for _,url := range urls {
		url := url
		t.Put(func(){
			_ = url
		})
	}
	t.Wait()
}

type TaskRunner struct {
	token chan bool
	wg sync.WaitGroup
}

func NewTaskRunner(concurrent int) *TaskRunner{
	return &TaskRunner{
		token:make(chan bool,concurrent),
	}
}

func(t *TaskRunner) Put(task func()) {
	t.token <- true
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		task()
		<- t.token
	}()
}

func (t *TaskRunner) Wait() {
	t.wg.Wait()
}