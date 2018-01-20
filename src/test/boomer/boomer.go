package main

import boomer "github.com/myzhan/boomer"

import (
	"log"
	"net"
	"net/http"
	"time"
)

/*
	forked from http://myzhan.github.io/2016/03/01/write-a-load-testing-tool-in-golang/
	added dns test cases.
*/

func now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func test_http() {
	/*
	   一个常规的 HTTP GET 操作，实际使用时，这里放业务自身的处理过程
	   只要将处理结果，通过 boomer 暴露出来的函数汇报就行了
	   请求成功，类似 Locust 的 events.request_success.fire
	   boomer.Events.Publish("request_success", type, name, 处理时间, 响应耗时)
	   请求失败，类似 Locust 的 events.request_failure.fire
	   boomer.Events.Publish("request_failure", type, name, 处理时间, 错误信息)
	*/
	startTime := now()
	resp, err := http.Get("http://localhost:8080/")
	defer resp.Body.Close()
	endTime := now()
	log.Println(float64(endTime - startTime))
	if err != nil {
		boomer.Events.Publish("request_failure", "demo", "http", 0.0, err.Error())
	} else {
		boomer.Events.Publish("request_success", "demo", "http", float64(endTime-startTime), resp.ContentLength)
	}
}

func test_dns() {
	/*
		一个简单的 DNS查询操作
		TODO: complicated dns query using https://github.com/miekg/exdns/blob/master/q/q.go
	*/
	startTime := now()
	addrs, err := net.LookupHost("www.baidu.com")
	endTime := now()

	log.Println("resolving www.baidu.com.. cost time: ", float64(endTime-startTime))
	if err != nil {
		boomer.Events.Publish("request_failure", "dns", "udp", 0.0, err.Error())
	} else {
		boomer.Events.Publish("request_success", "dns", "udp", float64(endTime-startTime), int64(len(addrs)))
	}
}

func main() {

	task := &boomer.Task{
		Name: "dns",
		// Weight 权重，和 Locust 的 task 权重类似，在有多个 task 的时候生效
		// FIXED: 之前误写为Weith
		Weight: 10,
		// Fn 类似于 Locust 的 task
		Fn: test_dns,
	}

	/*
	   通知 boomer 去执行自定义函数，支持多个
	   boomer.Run(task1, task2, task3)
	*/

	boomer.Run(task)

}
