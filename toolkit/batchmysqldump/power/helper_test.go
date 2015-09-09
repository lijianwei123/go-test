package power

import (
	"fmt"
	"testing"
)

//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/11.3.md  单元测试写法
func Test_System(t *testing.T) {
	if system("who") != "" {
		fmt.Println(System("who"))
		t.Log("system 测试通过")
	} else {
		t.Error("system 测试不通过")
	}
}
