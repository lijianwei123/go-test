package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var reqCount, reqConcurrent int
	var url string

	flag.IntVar(&reqCount, "n", 1, "-n requests")
	flag.IntVar(&reqConcurrent, "c", 1, "-c concurrency")
	flag.StringVar(&url, "url", "http://www.baidu.com", "-url url")
	flag.Parse()
	fmt.Printf("reqCount: %d, reqConcurrent: %d, url: %s\n", reqCount, reqConcurrent, url)

	times := reqCount / reqConcurrent
	chans := make(chan bool, reqCount)
	start := time.Now().UnixNano() / 1e6

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
			if n%100 == 0 {
				fmt.Println(n)
			}
		}
	}

	elaspedMs := time.Now().UnixNano()/1e6 - start
	fmt.Printf("elasped(ms): %d\n", elaspedMs)
	fmt.Printf("qps: %f\n", float64(reqCount)/float64(elaspedMs)*1000)
}
