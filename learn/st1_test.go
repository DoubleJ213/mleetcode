package learn

import (
	"fmt"
	"testing"
)

var s11 = []int{1, 1, 1}

func TestSt1(t *testing.T) {
	/*s := []int{5}
	fmt.Println(cap(s))
	s = append(s, 7)
	fmt.Println(cap(s))
	s = append(s, 9)
	fmt.Println(cap(s))
	x := append(s, 11)
	s = append(s, 10)
	y := append(s, 12)
	s = append(s, 13)
	fmt.Println(s, x, y)
	//https://golang.design/go-questions/slice/grow/
	//注意slice容量。
	*/

	//fmt.Println(mySqrt(48))
	/*
		s := []int{1, 2}
		s = append(s, 4, 5, 6)
		fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))
		s = append(s, 7, 8)
		fmt.Printf("len=%d, cap=%d \n", len(s), cap(s))*/
	/*	fmt.Println(uintptr(5))
		fmt.Println(^uintptr(0) >> 63)
		fmt.Println(^uintptr(0))
		fmt.Println(^uintptr(0)>>63)*/
	/*
		s := []int{1, 1, 1}
		//s := [3]int{1, 1, 1}
		//
		//for i := range s {
		//	s[i] += 1
		//}

		//maxIndex := len(s) - 1
		//for i, e := range s {
		//	if i == maxIndex {
		//		s[0] += e
		//	} else {
		//		s[i+1] += e
		//	}
		//}
		//
		//for i := 0; i < len(s); i++ {
		//	if i == maxIndex {
		//		s[0] += s[i]
		//	} else {
		//		s[i+1] += s[i]
		//	}
		//}

		f(s)
		fmt.Println(s)*/
	/*
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s1 := s
		s2 := s1[6:]
		s3 := s1[6:8]
		s1[2] = 100
		fmt.Println(s)
		fmt.Println(s1)
		fmt.Println(s2)
		s2 = append(s2, 200)
		s3 = append(s3, 300)
		fmt.Println(s)
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(s3)*/
	//fmt.Println(newS)

	//s1 := [3]int{1, 1, 1}
	//newS1 := myAppend1(s1)
	//fmt.Println(newS1)
	//
	//s = newS
	//
	//myAppendPtr(&s)
	//fmt.Println(s)
	/*a := make([]int, 3, 4)
	fmt.Println(a, len(a), cap(a))
	a = []int{1, 1, 1}
	fmt.Println(a, len(a), cap(a))
	b := a[1:2]
	c := a
	b = append(b, 100)
	c = append(c, 200)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b)
	fmt.Println(c, len(c), cap(c))
	aa := make([]int, 4, 5)
	aa[1] = 1
	aa[2] = 2
	fmt.Println(aa, len(aa), cap(aa))
	bb := aa[0 : len(aa)-1]
	fmt.Println(bb, len(bb), cap(bb))
	bb = append(bb, 4)
	fmt.Println(aa, len(aa), cap(aa))
	fmt.Println(bb, len(bb), cap(bb))*/

	//newS := myAppend(s11)
	////newS := myAppend2(s11)
	//
	//fmt.Println(s11)
	//fmt.Println(newS)
	//
	//newS1 := myAppend2()
	//fmt.Println(s11)
	//fmt.Println(newS1)
	//
	//s11 = newS
	//
	//myAppendPtr(&s11)
	//fmt.Println(s11)

	//fmt.Println("1 -- :", 7%3)
	//fmt.Println("-1 -- :", -7%3)

	golang := Go{}
	php := PHP{}

	//sayHello(golang)
	//sayHello(php)

	/*
		调用 sayHello() 函数时，传入了 golang, php 对象，
		它们并没有显式地声明实现了 IGreeting 类型，只是实现了接口所规定的 sayHello() 函数。
		实际上，编译器在调用 sayHello() 函数时，会隐式地将 golang, php 对象转换成 IGreeting 类型，
		这也是静态语言的类型检查功能。	顺带再提一下动态语言的特点：
		变量绑定的类型是不确定的，在运行期间才能确定 函数和方法可以接收任何类型的参数，
		且调用时不检查参数类型 不需要实现接口
		总结一下，鸭子类型是一种动态语言的风格，在这种风格中，一个对象有效的语义，
		不是由继承自特定的类或实现特定的接口，而是由它"当前方法和属性的集合"决定。
		Go 作为一种静态语言，通过接口实现了 鸭子类型，实际上是 Go 的编译器在其中作了隐匿的转换工作。
	*/
	//golang.sayHello()
	//php.sayHello()

	var i IGreeting
	i = golang
	sayHello(i)
	i = php
	sayHello(i)

	// qcrao 是值类型
	qcrao := person{age: 18}

	// 值类型 调用接收者也是值类型的方法
	fmt.Println(qcrao.howOld())

	// 值类型 调用接收者是指针类型的方法
	qcrao.growUp()
	//在Go语言中，当一个方法的接收者是指针类型时，如果该方法被一个值类型的变量调用，
	// Go会在后台自动将该值类型的变量的地址取出来，然后调用这个方法。
	// 这种自动取地址的行为是Go语言的一种语法糖，让代码编写更加方便。
	fmt.Println(qcrao.howOld())

	// ----------------------

	// stefno 是指针类型
	stefno := &person{age: 100}

	// 指针类型 调用接收者是值类型的方法
	fmt.Println(stefno.howOld())

	// 指针类型 调用接收者也是指针类型的方法
	stefno.growUp()
	fmt.Println(stefno.howOld())
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	//for _, i := range s {
	//	i++
	//}
	for i := range s {
		fmt.Println(i)
		s[i] += 1
	}
}
func myAppend(s11 []int) []int {
	// 这里 x 虽然改变了，但并不会影响外层函数的 s
	s11 = append(s11, 100)
	return s11
}
func myAppend2() []int {
	// 这里 x 虽然改变了，但并不会影响外层函数的 s
	s11 = append(s11, 100)
	return s11
}
func myAppend1(s [3]int) [3]int {
	//数组没有append方法
	//func append(slice []Type, elems ...Type) []Type
	//s = append(s, 100)
	return s
}
func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

type person struct {
	age int
}

func (p person) howOld() int {
	return p.age
}

func (p *person) growUp() {
	p.age += 1
}
