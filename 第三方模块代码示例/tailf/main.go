package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	filenames := "/golang/go_dev/Project/Exercise/tailf/log.txt"
	tails, err := tail.TailFile(filenames, tail.Config{
		// 文件被检测到变化，会重新打开。
		ReOpen:    false,
		// 文件被移走或重命名，tailf会读取变更后的新的文件。
		Follow:    true,
		// 文件异常，如读文件的程序中断，已经往kafka写了一些，下次再从头开始读就会重复。
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		// 文件非必须存在，比如新上线一台业务机器，日志信息没有及时打印输出，并不代表之后不会有日志信息产生。
		MustExist: false,
		// 对于请求量非常少的系统，读到最新位置没有信息可读了，此时就需要去不断查询。
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	// msg表示一行数据。
	var msg *tail.Line
	var ok bool
	for true {
		// 管带关闭，OK等于False。
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
