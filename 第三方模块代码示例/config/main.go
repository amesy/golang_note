package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
	"encoding/json"
)

var res = make(map[string]interface{})

func main() {
	filename := "/golang/go_dev/Project/Exercise/config/file_ini"
	conf, err := config.NewConfig("ini", filename)
	if err != nil {
		fmt.Println("New config file error, err: ", err)
		return
	}

	server_ip := conf.String("server::server_ip")
	res["server_ip"] = server_ip

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read config file port error, err:", err)
	}
	res["port"] = port

	log_path := conf.String("server::log_path")
	res["log_path"] = log_path

	fmt.Println(res)

	// encode.
	lst, err := json.Marshal(res)

	// decode.
	var j_lst interface{}

	json.Unmarshal([]byte(lst), &j_lst)

	fmt.Println("j_lst:", j_lst)

}