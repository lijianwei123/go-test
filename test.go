package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "test_koudaitong:nPMj9WWpZr4zNmjz@tcp(192.168.66.202:3306)/test_koudaitong?charset=utf8", 30)
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	var lists []orm.ParamsList

	num, err := o.Raw("show tables").ValuesList(&lists)
	if err == nil && num > 0 {
		fmt.Println(lists[0][0])
	}
}
