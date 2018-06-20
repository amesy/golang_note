package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
	"log"
)

func init() {
	// 自定义日志格式信息。
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	config := make(map[string]interface{})
	// 若日志文件不存在，linux下会自动创建。
	config["filename"] = "/golang/go_dev/Project/Exercise/logs/logs.log"
	// 定义日志级别。
	config["level"] = logs.LevelWarn

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("config file err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	logs.Debug("this is a debug log.")
	//log.Fatalln("fatal message")
	//log.Panicln("panic message")
	logs.Trace("this is a trace log.")
	logs.Warn("this is a warn log.")
}
