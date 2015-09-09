package work

import (
	"runtime"
)

func DoMonitor(config MysqlConfig) {
	//连接数据库
	threadInfos := make(map[int]MysqlThreadInfo)
	conn := MysqlConn{}

	for {
		//show processlist

		for _, threadInfo := range threadInfos {
			if threadInfo.CanKillQuery(config.MaxExecTime) {
				conn.KillQuery(threadInfo.ThreadId)
				/记录kill sql
			}
		}

		runtime.Gosched()
	}
}
