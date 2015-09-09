package main

import (
	"fmt"
	_ "github.com/lijianwei123/go-test/toolkit/batchmysqldump/power"
	"github.com/lijianwei123/go-test/toolkit/qs_bigile/work"
)

func main() {

	var search = "lijianwei"

	//上下文
	var ctxNum = 10

	var filePath = "/data/logs/pay.log"

	_, _, _ = search, ctxNum, filePath

}
