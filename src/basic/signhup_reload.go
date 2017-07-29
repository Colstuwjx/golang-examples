package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	hup := make(chan os.Signal)
	signal.Notify(hup, syscall.SIGHUP)

	go func() {
		for {
			select {
			// NOTE: 在mac上，执行`go run sighup_reload.go`
			// 将会发现两个相关进程
			// 501 27655  5328   0  9:44下午 ttys000    0:00.05 go run signhup_reload.go
			// 501 27661 27655   0  9:44下午 ttys000    0:00.00 /var/folders/qp/xygy8gm97lsdn6w7lv3pb2lc0000gn/T/go-build138388009/command-line-arguments/_obj/exe/signhup_reload
			// 发送信号给下面这个`PID=27661`的方能按照预期那样生效.
			case <-hup:
				// do reload
				// such as `reloadConfig` method
				// and safely replace the config instance.
				fmt.Println("received ", hup)
				fmt.Println("do reload here...")
				// default part should be blank...
				// default:
				// 	fmt.Println("no hup signal...")
				// 	time.sleep(1)
				// 	continue
			}
		}
	}()

	fmt.Println("some faked server started...")
	for i := 1; i <= 100; i++ {
		fmt.Println("I'm doing jobs...")
		time.Sleep(10 * time.Second)
	}
}
