package main

import (
	"fmt"
	"github.com/mikitu/golang-workers-pool/src/pool"
	"github.com/mikitu/golang-workers-pool/src/tasks"
)

func main() {
	pool1 := wpool.NewPool(5)

	pool1.Exec(tasks.ExampleTask("foo"))
	pool1.Exec(tasks.ExampleTask("bar"))

	pool1.Resize(3)

	pool1.Resize(6)

	for i := 0; i < 20; i++ {
		pool1.Exec(tasks.ExampleTask(fmt.Sprintf("additional_%d", i+1)))
	}

	pool1.Close()

	pool1.Wait()
}
