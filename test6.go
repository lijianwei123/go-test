package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("moring")
		case t.Hour() < 17:
			fmt.Println("aftermoon")
		default:
			fmt.Println("evening")
	}
}
