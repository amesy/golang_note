package main

import (
	"fmt"
	"sync"
	"time"
	"sync/atomic"
)

var mapMutex *sync.Mutex

type Ban struct {
	visitIPs map[string]struct{}
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]struct{})}
}

// 判断ip是否存在
func (o *Ban) visit(ip string) bool {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = struct{}{}
	go o.invalidAfter3Min(ip)
	return false
}

// 3分钟后ip失效，从map中移除，之后ip再次访问即可正常。
func (o *Ban) invalidAfter3Min(ip string) {
	time.Sleep(3 * time.Minute)
	mapMutex.Lock()
	visitIPs := o.visitIPs
	delete(visitIPs, ip)
	o.visitIPs = visitIPs
	mapMutex.Unlock()
}

func main() {
	mapMutex = new(sync.Mutex)
	var success int64
	ban := NewBan()
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			// ipEnd := j
			go func(j int) {
				defer wg.Done()
				// fmt.Println(j)
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	fmt.Println("success:", success)
}
