package power

import (
	_ "fmt"
)

var defaultConfig = DbConfig{"127.0.0.1", "root", 3306, "test", ""}

func GetDefaultDbConfig() DbConfig {
	return defaultConfig
}

type DbConfig struct {
	host     string
	passwd   string
	port     uint16
	database string
	table    string
}

type AlarmInfo struct {
	Title   string
	Type    string
	Content string
}
