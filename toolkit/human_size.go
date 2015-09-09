package main

import (
	"flag"
	"fmt"
	"github.com/lijianwei123/go-test/toolkit/qs_bigfile/work"
	"os"
)

//在mac下查看文件大小   KB   MB  GB  TB
func human(size int64) string {

	var humanStr string = ""

	if size < 1024 {
		humanStr = fmt.Sprintf("%d%s", size, "B")
	} else if size < 1024*1024 {
		humanStr = fmt.Sprintf("%f%s", size/1024, "KB")
	} else if size < 1024*1024*1024 {
		humanStr = fmt.Sprintf("%f%s", size/(1024*1024), "MB")
	} else if size < 1024*1024*1024*1024 {
		humanStr = fmt.Sprintf("%f%s", size/1024/1024/1024, "GB")
	}

	return humanStr
}

func main() {
	flagSet := flag.NewFlagSet("human", flag.ExitOnError)
	filePath := flagSet.String("f", "", "file path")
	flagSet.Parse(os.Args[1:])

	if *filePath == "" {
		panic("file path empty")
	}

	fmt.Println("filePath = " + *filePath)

	fp := work.NewFile(*filePath)

	if !fp.IsExist() {
		fmt.Println(fp.GetFilePath() + " don't exist!")
		os.Exit(-1)
	}

	if !fp.IsFile() {
		fmt.Println(fp.GetFilePath() + " is not file")
		os.Exit(-1)
	}

	fileSize := fp.GetFileSize()

	fmt.Printf("%s size = %s\n", *filePath, human(fileSize))
}
