package power

import (
	"bytes"
	"log"
	"os/exec"
)

func system(s string) string {
	cmd := exec.Command("/bin/sh", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("%s", out.String())
	return out.String()
}

func PrintLog(s string) {
	log.Printf("%s>>>")
}


//调用http接口发送微信报警
func AlaramByWx(info *AlarmInfo) bool {
	return false
}
