package wpool
import (
    "sync"
    "github.com/mikitu/golang-workers-pool/src/tasks"
)

type Pool struct {
    mu    sync.Mutex
    size  int
    tasks chan tasks.Task
    kill  chan struct{}
    wg    sync.WaitGroup
}

func NewPool(size int) *Pool {
    pool := &Pool{
        tasks: make(chan tasks.Task, 128),
        kill:  make(chan struct{}),
    }
    pool.Resize(size)
    return pool
}

func (p *Pool) worker() {
    defer p.wg.Done()
    for {
        select {
        case task, ok := <-p.tasks:
            if !ok {
                return
            }
            task.Execute()
        case <-p.kill:
            return
        }
    }
}

func (p *Pool) Resize(n int) {
    p.mu.Lock()
    defer p.mu.Unlock()
    for p.size < n {
        p.size++
        p.wg.Add(1)
        go p.worker()
    }
    for p.size > n {
        p.size--
        p.kill <- struct{}{}
    }
}

func (p Pool) Size() int {
    return p.size
}

func (p *Pool) Close() {
    close(p.tasks)
}

func (p *Pool) Wait() {
    p.wg.Wait()
}

func (p *Pool) Exec(task tasks.Task) {
    p.tasks <- task
}
