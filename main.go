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
	/*	diff := getDiff(9, 8, 19, 14)
		diff += getDiff(9, 23, 18, 18)
		diff += getDiff(9, 17, 18, 50)
		diff += getDiff(9, 12, 17, 57)
		diff += getDiff(8, 29, 18, 0)
		fmt.Printf("比正常工时多(负数为少) %v \n", diff)*/

	numbers1 := [5]int{1, 2, 3, 4, 5} //切片
	for i := range numbers1 {
		if i == 5 {
			numbers1[i] &= i
			//numbers1[i] = numbers1[i] | i
			//	按位或  4|3 100 ｜ 011 = 111  111的十进制7
		}
	}
	fmt.Println(numbers1)

	numbers2 := []int{1, 2, 3, 4, 5} //数组
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] &= e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	//	[5 3 5 7 9 11]
	/*
		numbers3 := []int{1, 2, 3, 4, 5, 6} //数组
		maxIndex3 := len(numbers3) - 1
		for i, e := range numbers3 {
			if i == maxIndex3 {
				numbers3[0] |= e
			} else {
				numbers3[i+1] += e
			}
		}
		fmt.Println(numbers3)*/
	//[5 3 6 10 15 21]

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
