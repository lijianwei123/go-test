package main

import (
	"fmt"
	_ "io/ioutil"
	_ "net/http"
	_ "os"
)

func test() {
	var ch = make(chan int)

	go func() {
		fmt.Println("lijianwei")
		ch <- 1
	}()
	<-ch
	fmt.Println("weiyanping")
}

func main() {
	test()
}
