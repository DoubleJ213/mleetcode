package test

import (
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

type Person struct {
	name string
	age  *int
}

func TestST(t *testing.T) {
	i := 19
	//p:=&Person{name:"张三",age:&i}
	p := Person{name: "张三", age: &i}
	fmt.Println(p)
	modify(p)
	fmt.Println(p)
}

func (p Person) String() string {
	return "姓名为：" + p.name + ",年龄为：" + strconv.Itoa(*p.age)
}

func modify(p Person) {
	//func modify(p *Person){
	p.name = "李四"
	*p.age = 20
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	fmt.Printf("head: %p\n", &head)
	var res *ListNode
	for head != nil {
		head.Next, res, head = res, head, head.Next
	}
	return res
}

func TestReverseNode(t *testing.T) {
	res := &ListNode{0, nil}
	temp := res
	for i := 1; i < 6; i++ {
		temp.Next = &ListNode{i, nil}
		temp = temp.Next
	}
	fmt.Printf("res.Next: %p\n", &res.Next)
	reverseList(res.Next)
}

func TestST2(t *testing.T) {
	i := 10
	ip := &i
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	modify2(ip)
	fmt.Println("int值被修改了，新值为:", i)

	var liLei = '('
	var liLei1 = "("
	findType(liLei)
	findType(liLei1)

	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))
}

func findType(x interface{}) {
	//lilei 字符串类型
	switch a := x.(type) {
	case byte:
		fmt.Printf("lilei is byte %v %v \n", a, x)
	case rune:
		fmt.Printf("lilei is rune %v %v \n", a, x)
	case int:
		fmt.Printf("lilei is int %v %v \n", a, x)
	case nil:
		fmt.Printf("lilei is nil %v %v \n", a, x)
	case string:
		fmt.Printf("lilei is string %v %v \n", a, x)
	default:
		fmt.Printf("lilei is default %v %v \n", a, x)
	}
}

func modify2(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 1
}
