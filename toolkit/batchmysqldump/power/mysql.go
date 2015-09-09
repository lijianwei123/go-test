package power

import (
	"fmt"
)

type mysql interface {
	connect(dbConfig Config) int
}

func Test() {
	name := "lijianwei"
	fmt.Printf("my name is %s\n", name)
}

//获取所有表
func GetTables(dbName string) {

}
