package main

import (
	_ "github.com/bitly/go-nsq"
	"fmt"
	"github.com/lijianwei123/go-test/toolkit/batchmysqldump/power"
)

//获取topic下的channel
func GetTopicChannel(topic string) []string {
	return []string{}
}

//通过http请求查询到相关信息
func GetChannelStatInfo(channel string) map[string]interface{} {
	statInfo = make(map[string]interface{})
	
	return statInfo
}

func main() {
	topic := "binlog_pay_order_finished"
	channel := "binlog_pay_order_finished"
	
	alarmNum := 1000
	alarmInfo := power.AlarmInfo{title: "nsq待消费数量过多", type: "nsq_need_consume_num", content: alarmNum}
	
	
	statInfo := GetChannelStatInfo(channel)
	if statInfo.num > alarmNum {
		power.AlaramByWx(alarmInfo)
	}
}



