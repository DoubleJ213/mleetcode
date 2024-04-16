package learn

import (
	"fmt"
	"testing"
	"time"
)

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello, world"
	done <- true
}

func aGoroutine1() {
	msg = "hello, world"
	fmt.Println("1", msg)
	time.Sleep(time.Second * 10)
	<-done
	time.Sleep(time.Second * 1)
	//有概率打印不出来
	fmt.Println("2", msg)
}

func TestSt4(t *testing.T) {
	//go aGoroutine()
	//<-done
	//println(msg)

	go aGoroutine1()
	println("3", msg)
	done <- true
	println(msg)
}


func worker() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <- ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}