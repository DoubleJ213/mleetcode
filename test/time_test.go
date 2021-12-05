package test

import (
	"testing"
	"time"
)

func report(startTime, endTime time.Time, ts []time.Time) []bool {
	h := int(endTime.Sub(startTime).Hours())
	if endTime.Minute() > 0 {
		h = h + 1
	}
	res := make([]bool, h)
	//初始化的res，有h个元素，索引h-1
	if len(ts) > 0 {
		interval(startTime, endTime, ts, res, 0, h-1)
	}
	return res
}

func interval(startTime, endTime time.Time, ts []time.Time, res []bool, left, right int) []bool {
	// 表示开关机状态
	status := false
	if startTime.Before(ts[0]) {
		/* 找ts数组对应的第一个整点，对应的res的第几个元素
		因为需要确认从res哪个索引开始修改电源状态*/

		//	逐个比较，这里可以简单的就for循环
		//	这个for循环可以搞个方法，因为下面也用了应该是差不多逻辑，不同的是从res哪个值开始改
		//	可以考虑搞个递归的方法 递归的条件就是，开始终止时间，和未判断的ts，res的值，以及被写过的索引
		//	每次就写这个res
		//	递归终止条件就是res最后一个元素都被写过了

		for i := 0; i < len(ts); i++ {
			//这里判断res对应的整点对应状态，就改res数组就行了

			//交替变化的开机关机状态
			status = changeStatus(status)
		}

	} else {
		/*
			当需要的打点时间小于ts的第一个元素时间，表示打点的开始时间足够早
			res的值从第一个就可以被修改
		*/

		for i := 0; i < len(ts); i++ {
			//这里判断res对应的整点对应状态，就改res数组就行了

			//交替变化的开机关机状态
			status = changeStatus(status)
		}
	}

	/*这里有个优化，endTime和ts[len(ts)-1]不比较。
	因为上面我们已经把起始确定好，且把ts数组遍历完了。
	当ts足够长打点时间还在endTime之后，考虑剪枝，当res最后一个元素都被修改之后，可以直接return
	当ts长度不够长，我们在之前从ts的第一个整点开始改，改到所有ts包含的整点都改完了，后面没有打点时间的全是false就行

	确认下是否可以这样处理！*/

	return res
}

func changeStatus(bool2 bool) bool {
	if bool2 {
		return false
	}
	return true
}

func TestReportA(t *testing.T) {
	startTime, _ := time.Parse(time.RFC3339, "2021-10-09T22:00:09Z")
	t0, _ := time.Parse(time.RFC3339, "2021-10-09T22:00:10Z")
	t6, _ := time.Parse(time.RFC3339, "2021-10-10T20:55:00Z")
	endTime, _ := time.Parse(time.RFC3339, "2021-10-11T03:59:59Z")
	ts := []time.Time{t0, t6}
	report(startTime, endTime, ts)
}
