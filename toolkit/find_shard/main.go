//add by laoli

package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"strconv"
	"time"
	"sync"
	"bufio"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	
	"github.com/c4pt0r/cfg"
	
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Conf struct {
	host     string
	port	 int
	user	 string
	pwd      string
	db		 string
}

func LoadConf(configFile string) (*Conf, error) {
	cobarConf := &Conf{}
	
	conf := cfg.NewCfg(configFile)
	if err := conf.Load(); err != nil {
		log.Fatalf("load config error: %s", err)
	}
	

	cobarConf.host, _ 	=	conf.ReadString("cobar_host", "127.0.0.1")
	cobarConf.port, _ 	=	conf.ReadInt("cobar_port", 8066)
	cobarConf.user, _ 	=	conf.ReadString("cobar_user", "test")
	cobarConf.pwd,  _ 	= 	conf.ReadString("cobar_pwd", "test")
	cobarConf.db,   _ 	= 	conf.ReadString("cobar_db", "test")
	
	return cobarConf, nil
}

func main() {
	var flagSet = flag.NewFlagSet("find_shard", flag.ExitOnError)
	var configFile = flagSet.String("config", "", "config file")
	
	flagSet.Parse(os.Args[1:])
	
	if *configFile == "" {
		log.Fatalf("Error: --config empty")
	}
	
	fmt.Printf("configFile:%s\n", *configFile)
	
	//读取配置
	cobarConf, _ := LoadConf(*configFile)
	
	//_ = cobarConf
	
	//fmt.Println(cobarConf)
	
	dsn := cobarConf.user + ":" + cobarConf.pwd + "@tcp(" + cobarConf.host + ":" + strconv.Itoa(cobarConf.port) + ")/" + cobarConf.db
	
	fmt.Print(dsn + "\n")
	
	//连接  https://github.com/go-sql-driver/mysql/wiki/Examples
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}
	
	//mysql 查询示例   http://www.cnblogs.com/hupengcool/p/4143238.html
	result, err := db.Query("show tables")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	
	columns, _ :=  result.Columns()
	
	fmt.Printf("len(columns):%d\n", len(columns))
	
	fmt.Println(columns)
	
	fmt.Println("===========")
	
	scanArgs := make([] interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	
	_ = scanArgs
	
	var table string
	
	_ = table
	
	for result.Next() {
		if err := result.Scan(scanArgs...); err != nil {
			panic(err.Error())
		}
		
		//fmt.Println(string(values[0].([]byte)))
	}
	
	var cobar_shard_template string = "/*!cobar: $dataNodeId=%d.0, $table='%s'*/"
	
	//无缓存 channel
	var id int = 0
	var ch = make(chan int, 1)
	var waitGroup sync.WaitGroup
	
	//go select http://www.sharejs.com/codes/go/4415
	//go 并发  http://studygolang.com/articles/2027
	
	//注册信号处理函数
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("ctrl-c or SIGTERM found, exit")
		os.Exit(-1)
	}()
	
	var find_shard = func(find_sql string, table string) {
		for i := 1; i <= 1024; i++ {
			if i % 10 == 0 {
				//堵着
				waitGroup.Wait()
				
				if id > 0 {
					//防止有错误发生
					select {
						case shard := <-ch:
							fmt.Println("你家的娃在" + strconv.Itoa(shard - 1) + "分片里")
						case <-time.After(time.Duration(3) * time.Second):
							fmt.Println("这是要弄啥呢")
						default:
							fmt.Println("我啥也不干")
					}
					break
				}
			}
			waitGroup.Add(1)
			
			go func(shard int) {
				defer waitGroup.Done()
				if id > 0 {
					//已经找到了
					return;
				}
				sql := fmt.Sprintf(cobar_shard_template + find_sql, shard-1, table)
	
				row := db.QueryRow(sql)
				
				err := row.Scan(&id)
				
				if err == nil && id > 0 {
					ch<-shard
				}
			}(i)
		}
	}
	
	reader := bufio.NewReader(os.Stdin)
	for {
		usage := 
`input select sql
like  select id from pay_order where order_no = '386920340142344'`
		
		fmt.Println(usage)
		
		input, _ := reader.ReadBytes('\n')
		input_string := string(input[0:len(input)-1])
		input_string = strings.TrimSpace(input_string)
		
		fmt.Println("你的输入为" + input_string)
		
		if input_string != "" {
			//找出table
			reg := regexp.MustCompile(`from\s+(\w+)\s+`)
			if table := reg.FindStringSubmatch(input_string); table != nil {
				table_name := table[1]
				if  table_name == "" {
					fmt.Println("can not find table name")
				} else {
					find_shard(input_string, table_name)
				}
			}
		}
	}
	
	//go debug  使用gdb 
	//sudo brew install homebrew/dupes/gdb
	//处理不能调试的  http://www.beanmoon.com/2014/11/23/install_gdb_in_mac/
}
