package main

import (
	"github.com/lijianwei123/go-test/toolkit/batchmysqldump/power"

	"fmt"

	_ "os/exec"

	_ "bytes"

	_ "log"
)

func main() {

	defaultConfig := power.GetDefaultConfig()

	_ = defaultConfig

	//协程个数
	var c int = 10

	fmt.Printf("%d", c)

	//var channel chan string = make(chan string, c)
}
