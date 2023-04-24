package review

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Long struct {
	Long string //这里面“Long”首字母必须大写，否则下面的ss打印出来就是“{ }”取不到数据。大坑 大坑
}

type Speaker interface {
	Hello()
}

type Speak struct {
	Name string
	Age  int
}

func (this *Speak) Hello() {
	fmt.Println("hello my name is", this.Name)
}

func TestCp(t *testing.T) {
	src := make([]int, 0)
	src = append(src, 1)
	src = append(src, 2)
	src = append(src, 3)
	src = append(src, 4)
	fmt.Print(src)
	tmp := make([]int, 0)
	//slice会在append自动扩容，但是初始长度为0
	// Copy returns the number of elements copied,
	// which will be the minimum of len(src) and len(dst).
	copy(tmp, src)
	tmp1 := make([]int, len(src))
	copy(tmp1, src)
	//将src的元素追加到[]int{}之后
	tmp2 := append([]int{}, src...)
	_ = tmp2
	tmp3 := append([]int(nil), src...)
	_ = tmp3
	fmt.Print("done")

	var s Speaker = &Speak{"wss", 10}
	//s = TestUser{"wss", 10}
	s.Hello()

	var ss Long
	str := `{"long":"4462"}`
	json.Unmarshal([]byte(str), &ss)
	fmt.Println(ss)

	//interface 不是单纯的值，而是分为类型和值。
	//所以传统认知的此 nil 并非彼 nil，
	//必须得类型和值同时都为 nil 的情况下，interface 的 nil 判断才会为 true。
	//所以传统认知的此 nil 并非彼 nil，必须得类型和值同时都为 nil 的情况下，interface 的 nil 判断才会为 true。 Go interface 是 Go 语言中最常用的类型之一，大家用惯了 if err != nil 就很容易顺手就踩进去了

	var v interface{}
	v = (*int)(nil)
	fmt.Println(v == nil)

	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)

	in = data
	fmt.Println(in, in == nil)

}

// LimitRangesGetter has a method to return a LimitRangeInterface.
// A group's client should implement this interface.
type LimitRangesGetter interface {
	LimitRanges(namespace string)
	LimitRangeInterface
}

// LimitRangeInterface has methods to work with LimitRange resources.
type LimitRangeInterface interface {
	Create()
	Update() LimitRangesGetter
}

// limitRanges implements LimitRangeInterface
type limitRanges struct {
	// interface struct 大小写潜规则
	ns string
}
