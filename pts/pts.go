package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

func main() {
	fmt.Println("fuck you?")

	reqConcurrent := 100
	reqCount := 1000
	times := reqCount / reqConcurrent
	url := "http://www.baidu.com"

	fmt.Println(reflect.TypeOf(times))

	chans := make(chan bool, reqCount)

	start := time.Now().Unix()

	for c := 0; c < reqConcurrent; c++ {
		go func() {

			for t := 0; t < times; t++ {
				resp, err := http.Get(url)
				defer resp.Body.Close()
				if err == nil && resp.StatusCode == http.StatusOK {
					ioutil.ReadAll(resp.Body)
				}

				chans <- true
			}
		}()
	}


	for n := 0; n < reqCount; n++ {
		select {
		case <-chans:
			fmt.Println(n)
		}
	}

	elasped := time.Now().Unix() - start
	fmt.Printf("qps: %d", reqCount/int(elasped))

}
