package work

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host        string
	Port        uint16
	User        string
	Passwd      string
	Db          string
	MaxExecTime uint16
}

func NewMysqlConfig() MysqlConfig {
	config := MysqlConfig{}
	config.Host = "127.0.0.1"
	config.Port = 3306
	config.User = "root"
	config.Passwd = "root"
	config.Db = ""
	config.MaxExecTime = 3

	return config
}

type ConnErr struct {
	errMsg string
}

func (errPtr *ConnErr) Error() string {
	return errPtr.errMsg
}

type MysqlConn struct {
	ConnName string
}

func NewMysqlConn(config MysqlConfig) (MysqlConn, error) {
	if config.Db == "" {
		return nil, &ConnErr{"db配置为空"}
	}

	configName := fmt.sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.User, config.Passwd, config.Host, config.Port, config.Db)

	orm.RegisterDataBase(config.Db, "mysql", configName, 30)

	return MysqlConn{config.Db}, nil
}

func (conn MysqlConn) KillQuery(threadId int) bool {
	o := orm.NewOrm()
	o.Using(conn.ConnName)

	sql := fmt.Sprintf("kill %d", threadId)
	o.Raw(sql).Exec()

	return true
}

type MysqlThreadInfo struct {
	ThreadId uint32
	User     string
	Host     string
	Db       string
	Command  string
	Time     uint16
	State    string
	Info     string //sql
}

func NewMysqlThreadInfo() MysqlThreadInfo {
	var threadInfo MysqlThreadInfo

	return threadInfo
}

func (threadInfo MysqlThreadInfo) CanKillQuery(maxExecTIme uint16) bool {
	return threadInfo.Command == "query" && threadInfo.Time >= maxExecTIme
}
