package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func main() {
	/*	var goOn = true
		var times int
		var diff time.Duration
		for goOn {
			times++
			diff += calculate(times)
			fmt.Printf("是否继续输入(是:t 不是:f)\n")
			fmt.Scanln(&goOn)
		}*/
	diff := getDiff(8, 33, 19, 04)
	diff += getDiff(9, 21, 19, 21)
	diff += getDiff(9, 29, 18, 50)
	diff += getDiff(9, 13, 19, 16)
	diff += getDiff(9, 20, 17, 30)
	//diff += getDiff(13, 30, 18, 30)
	fmt.Printf("比正常工时多(负数为少) %v \n", diff)
}

func getInputTime(
	nums int, word, word2 string) int {
	var myTime string
	var mt int
	err := errors.New("开始输入")
	for err != nil {
		fmt.Printf("请输入第%d天%s时间--%s:", nums, word, word2)
		fmt.Scanln(&myTime)
		mt, err = strconv.Atoi(myTime)
		if err != nil {
			fmt.Printf("输入异常，重新输入\n")
		}
	}
	return mt
}

func calculate(nums int) time.Duration {
	startH := getInputTime(nums, "上班", "小时")
	startM := getInputTime(nums, "上班", "分钟")
	endH := getInputTime(nums, "下班", "小时")
	endM := getInputTime(nums, "下班", "分钟")
	fmt.Printf("您输入的第%d天 上班时间%d:%d 下班 %d:%d \n", nums, startH, startM, endH, endM)
	return getTodayDiff(startH, startM, endH, endM)
}

func getDiff(startH, startM, endH, endM int) time.Duration {
	return getTodayDiff(startH, startM, endH, endM)
}

const (
	year        = 2023
	month       = time.Month(8)
	day         = 1
	defaultTime = time.Duration(34200000000000)
)

func getTodayDiff(startH, startM, endH, endM int) time.Duration {
	todayStart := time.Date(year, month, day, startH, startM, 0, 0, time.UTC)
	todayEnd := time.Date(year, month, day, endH, endM, 0, 0, time.UTC)
	return todayEnd.Sub(todayStart) - defaultTime
}
