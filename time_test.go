package mleetcode

import (
	"fmt"
	"testing"
	"time"
)

/*
这是一个生产中遇到的问题
题目描述:
现有一监控服务负责监控组件状态的变化，每当组件的状态发生变化，监控服务都会在数组中记录当前时间。
被监控的组件只有：false、true两种状态，且状态变化遵循false变成true，或者true变成false的规律，
不存在false之后又是false，或者true之后又是true的情况。
监控输出数据如[
2021-10-09T20:10:10,
2021-10-09T20:20:20,
2021-10-09T22:30:20,
2021-10-09T23:41:20,
2021-10-09T23:55:00,
2021-10-10T00:10:00
]

求一个时间范围比如【2021-10-09T19:00:00~2021-10-10T03:59:59】对应每个整点被监控组件的状态
输出数组[false false true true false false true true true]

补充说明：
1.监控时间数组，第一个时间对应的值为false
2.给定的范围为闭区间
3.给定的范围可以没有监控数据。
如果该整点前没有数据可以认为状态为false，如果该整点前有打点数据，按照前一个监控数据的状态来
*/
func changeStatus(bool2 bool) bool {
	//交替变化的开机关机状态
	if bool2 {
		return false
	}
	return true
}

func timeMod(t time.Time) time.Time {
	//向上取整，只找到了个向下取整的方法，简单处理下
	t = t.Truncate(time.Hour * 1)
	res, _ := time.ParseDuration("1h")
	return t.Add(res)
}

func timeAddHours(delta int, t time.Time) time.Time {
	addTime, _ := time.ParseDuration(fmt.Sprintf("%vh", delta))
	return t.Add(addTime)
}

func interval(startTime time.Time, ts []time.Time, res []bool) []bool {
	// 表示开关机状态
	status := false
	// 准备修改res的状态
	resStatus := false
	left := 0
	for i := 0; i < len(ts); i++ {
		// 这里判断res对应的整点对应状态，就改res数组就行了
		modTs := timeMod(ts[i])
		if modTs.Before(startTime) || modTs.Equal(startTime) {
			//ts0取整后早于startTime的情况，打点数据可以覆盖res对应的返回值
			resStatus = status
			if i+1 > len(ts)-1 || modTs != timeMod(ts[i+1]) {
				//与后一个值比较，如果不在同一个小时内，更新时间
				// i+1 > len(ts)-1 遍历到ts最后一个元素的时候，直接修改res，以及后面所有的res
				//这个比较可以考虑优化，直接字符比较啥的，这样转换效率比较低
				if resStatus {
					res[left] = resStatus
				}
				left += 1
				startTime = timeAddHours(1, startTime)
				if i+1 > len(ts)-1 {
					for j := left; j < len(res); j++ {
						if resStatus {
							res[j] = resStatus
						}
					}
					break
				}
			} else {
				fmt.Print("同一个小时内，等待下次循环\n")
			}
			status = changeStatus(status)
		} else {
			// ts0取整后晚于startTime的情况
			// 算出晚几个小时，即算出最终res数组，缺少几个值，缺少的值不能都是false，而是前一个时间戳对应的status
			delta := int(modTs.Sub(startTime).Hours())
			for j := left; j < left+delta; j++ {
				if resStatus {
					res[j] = resStatus
				}
			}
			left += delta
			startTime = timeAddHours(delta, startTime)
			//外层索引有减1的情况，会重复遍历某些元素，还是可以认为复杂度是O(N)
			//但内层有一个分支有for循环，虽然逻辑简单，只为了对res进行赋值，但是也有复杂度
			//整体来看，这样写的时间复杂度，直觉上说还是O(N*M)的o(n2)（n平方）
			//问题大大的，必须优化
			i--
		}
	}

	return res
}

func intervalB(startTime time.Time, ts []time.Time, res []bool) []bool {
	// 记录整点的状态
	resStatusMap := make(map[time.Time]bool)
	// 表示开关机状态
	status := false
	for i := 0; i < len(ts); i++ {
		resTime := timeMod(ts[i])
		resStatusMap[resTime] = status
		status = changeStatus(status)
	}
	for j := 0; j < len(res); j++ {
		if value, ok := resStatusMap[startTime]; ok {
			//  时间能找到
			res[j] = value
		} else {
			//	时间找不到的情况，用前一个时间的状态
			//  可以有剪枝操作，但是复杂度也不高了，不改了就
			if j == 0 {
				res[j] = false
			} else {
				res[j] = res[j-1]
			}
		}
		startTime = timeAddHours(1, startTime)
	}
	return res
}

func report(startTime, endTime time.Time, ts []time.Time) []bool {
	// 为了方便，可以startTime取整之后再传进来，也可以自己处理，因为本身是要整点的信息上报，这个时间不是整数也没有意义
	// 这里的startTime，endTime的意义是第一个整点和最后一个整点
	h := int(endTime.Sub(startTime).Hours())
	if endTime.Minute() > 0 {
		h = h + 1
	}
	res := make([]bool, h)
	//初始化的res，有h个元素，索引最大的h-1
	// 下面所有代码都是第一个打点数据为关机数据，因为初始化数组全false，如果不是再判断下
	if len(ts) > 0 {
		//interval(startTime, ts, res)
		intervalB(startTime, ts, res)
	}
	return res
}

func TestReportA(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339, "2021-10-09T19:00:00Z")
	t0, _ := time.Parse(time.RFC3339, "2021-10-09T20:10:10Z") // off
	t1, _ := time.Parse(time.RFC3339, "2021-10-09T20:20:20Z") // on
	t2, _ := time.Parse(time.RFC3339, "2021-10-09T22:30:20Z") // off
	t3, _ := time.Parse(time.RFC3339, "2021-10-09T23:41:20Z") // on
	t6, _ := time.Parse(time.RFC3339, "2021-10-09T23:55:00Z") // off
	t7, _ := time.Parse(time.RFC3339, "2021-10-10T00:10:00Z") // on
	endTime, _ := time.Parse(time.RFC3339, "2021-10-10T03:59:59Z")
	ts := []time.Time{t0, t1, t2, t3, t6, t7}
	report(startTime, endTime, ts)
}

/*
头部的电源状态，如果没有监控数据，都按照false处理
尾部的电源状态，按照最后一个打点数据来处理的
比如开机，那如果没有打点数据，认为他一直开机，全是true
*/
