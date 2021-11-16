package main

import (
	"52lu/go-pprof/scenes"
	"fmt"
	"net/http"
	_ "net/http/pprof" // 导入pprof
)
func init() {
	// 开启http端口,用协程的方式监听，否则会阻塞
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			fmt.Println("pprof err:",err)
		}
	}()
}
func main()  {
	fmt.Println("运行测试...")
	ch := make(chan bool)
	//scenes.UseHeapDemo()
	//scenes.UseCpuDemo()
	scenes.UseGoroutineDemo()
	<-ch
}

