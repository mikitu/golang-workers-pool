package tasks

import "fmt"

type ExampleTask string

func (e ExampleTask) Execute() {
	fmt.Println("executing:", string(e))
}
