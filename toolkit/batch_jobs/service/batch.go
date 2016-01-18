package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "os"
	"strconv"
	"sync"
)

type BatchConfig struct {
	Start   int `开始位置`
	Limit   int
	Size    int    `总的页数`
	ReqTmpl string `请求url`
}

type Batch struct {
	wg sync.WaitGroup
	bc BatchConfig
}

func NewBatch(bc BatchConfig) *Batch {
	var wg sync.WaitGroup

	return &Batch{
		wg: wg,
		bc: bc,
	}
}

func (bcPtr *Batch) Run() {

	for i := 1; i <= bcPtr.bc.Size; i++ {

		offset := bcPtr.bc.Start + (i-1)*bcPtr.bc.Limit
		url := fmt.Sprintf(bcPtr.bc.ReqTmpl, strconv.Itoa(offset), strconv.Itoa(bcPtr.bc.Limit))

		bcPtr.wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			fmt.Printf("%s response %s\n", url, string(body))

			defer bcPtr.wg.Done()
		}(url)
	}

	bcPtr.wg.Wait()
}
