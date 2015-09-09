package main

import (
	"flag"
	"fmt"
	"github.com/larspensjo/config" //读取config.ini
	"github.com/lijianwei123/go-logger/logger"
	"github.com/lijianwei123/go-test/sqlkill/work"
	fileWork "github.com/lijianwei123/go-test/toolkit/qs_bigfile/work"
	_ "github.com/samuel/go-zookeeper/zk"
	"os"
	_ "path"
	"path/filepath"
	"runtime"
	_ "strconv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//读取配置选项
	flagSet := flag.NewFlagSet("sqlkill", flag.ExitOnError)
	configPath := flagSet.String("f", "", "config path")
	flagSet.Parse(os.Args[1:])

	if *configPath == "" {
		panic("config path empty")
	}

	configFp := fileWork.NewFile(*configPath)
	if !configFp.IsExist() || !configFp.IsFile() || configFp.GetExt() != ".ini" {
		panic("config path" + *configPath + " error ")
	}

	absConfigPath, err := filepath.Abs(*configPath)

	//解析config.ini
	cnfData, err := config.ReadDefault(absConfigPath)
	if err != nil {
		panic("parse " + *configPath + "error:" + err.Error())
	}

	logPath, _ := cnfData.String("logger", "default.path")
	absLogPath, err := filepath.Abs(logPath)
	if err != nil {
		panic("log path error" + err.Error())
	}

	logDir, logFile := filepath.Dir(absLogPath), filepath.Base(absLogPath)

	//日志配置
	logger.SetRollingDaily(logDir, logFile)
	logger.SetLevel(logger.DEBUG)
	defer logger.CloseLogFile()

	//测试配置
	config1 := work.MysqlConfig{}
	config1.Host = "192.168.66.202"
	config1.Port = 3306
	config1.User = "test_koudaitong"
	config1.Passwd = "nPMj9WWpZr4zNmjz"
	config1.Db = ""
	config1.MaxExecTime = 1

	//读取配置
	killConfig := make(map[string]work.MysqlConfig)

	key1 := fmt.Sprintf("%s:%d", config1.Host, config1.Port)
	killConfig[key1] = config1

	for _, mysqlConfig := range killConfig {
		go work.DoMonitor(mysqlConfig)
	}

}
