package work

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	_ "log"
	"os"
	"os/exec"
)

func Pull() bool {
	gitWorkDir := beego.AppConfig.String("git_work_dir")
	os.Chdir(gitWorkDir)

	branchName, _ := system(fmt.Sprintf("echo `git rev-parse --abbrev-ref HEAD`"))
	cmd := fmt.Sprintf("git pull origin %s", branchName)
	output, _ := system(cmd)
	fmt.Println(output)

	return true
}

func system(s string) (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
