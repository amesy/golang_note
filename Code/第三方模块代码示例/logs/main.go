package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	config := make(map[string]interface{})
	config["filename"] = "/golang/go_dev/Project/Exercise/logs/logs.log"
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
