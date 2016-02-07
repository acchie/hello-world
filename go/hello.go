package main

import (
	"fmt"
	"time"
)

func trace(s string) string {
 fmt.Println("entering:", s)
return s
}
func un(s string) { fmt.Println("leaving:", s) }


func sub(ch chan bool) {
	defer un(trace("sub"))

	time.Sleep(3 *time.Second)
	fmt.Println("World.")
	ch <- true
}

func main() {
	defer un(trace("main"))

	fmt.Println("Hello, ")
	ch1 := make(chan bool)
	go sub(ch1)
	<- ch1
}
