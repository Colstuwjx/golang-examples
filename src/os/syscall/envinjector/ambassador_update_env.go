package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// demo for dynamically update env!
// #1 ambassador mode
func updateEnv() {
	random := rand.Intn(100)

	if random > 50 {
		os.Setenv("FOO", "BAR")
	} else {
		os.Setenv("HELLO", "Jacky")
	}
}

func printEnv() {
	fmt.Printf("Currently, we have: FOO=%s, HELLO=%s\n", os.Getenv("FOO"), os.Getenv("HELLO"))
}

func bindReload() {
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
				fmt.Println("received ", hup)
				updateEnv()
			}
		}
	}()
}

func main() {
	bindReload()

	for {
		printEnv()
		time.Sleep(10 * time.Second)
	}
}
