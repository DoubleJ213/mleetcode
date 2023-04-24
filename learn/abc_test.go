package learn

import (
	"fmt"
	"testing"
)

type Person struct {
	id   int
	name string
}

func TestAbc(t *testing.T) {
	p1 := Person{
		id:   1,
		name: "lisi",
	}
	p2 := &Person{
		id:   2,
		name: "wang",
	}
	var p3 Person = Person{
		3, "lap",
	}
	var p4 *Person = &Person{
		4, "xing",
	}
	chang(p1)                   //{20 lisi}
	p1.chang()                  //{1 lisi} {30 lisi}
	p1.changaddr()              //&{1 lisi} &{20 lisi}
	fmt.Println("p1", p1)       //p1 {1 lisi}  {20 lisi}
	fmt.Println("&p1", &p1)     //&p1 &{1 lisi}
	fmt.Println("*&p1", *&p1)   //*&p1 {1 lisi}    *&可以相互抵消
	fmt.Println("&*&p1", &*&p1) //&*&p1 &{1 lisi}
	fmt.Println("*p1 null")     //由于p1不是指针类型，所以无法用*获取指针的值
	fmt.Println("p3", p3)       //p3 {3 lap}
	fmt.Println("&p3", &p3)     //&p3 &{3 lap}
	fmt.Println("*p3 null")
	changaddr(p2)             //&{20 wang}
	p2.changaddr()            //&{20 wang} &{20 wang}
	p2.chang()                // {20 wang} {30 wang}
	fmt.Println("p2", p2)     //p2 &{2 wang}    &{20 wang}
	fmt.Println("&p2", &p2)   //&p2 0xc000006028 地址的地址
	fmt.Println("*&p2", *&p2) //*&p2 &{20 wang} 地址的地址所存的值
	fmt.Println("*p2", *p2)   //*p2 {2 wang} *取指针对象的值
	fmt.Println("p4", p4)     //p4 &{4 xing}
	fmt.Println("&p4", &p4)   //&p4 0xc000006030
	fmt.Println("*p4", *p4)   //*p4 {4 xing}
}

//传入的是一个拷贝的对象
func chang(person Person) {
	person.id = 20
	fmt.Println(person)
}

//传入的是一个指向传入对象的地址，会直接改变底层数值
func changaddr(person *Person) {
	person.id = 20
	fmt.Println(person)
}

// 都可当做对象调用，但只是拷贝了一个对象，不对原生值进行改变
func (person Person) chang() {
	fmt.Println("change")
	fmt.Println(person)
	person.id = 30
	fmt.Println(person)
}

// 都可当做对象调用，指向该对象地址，会对原生值进行改变
func (person *Person) changaddr() {
	fmt.Println("changaddr")
	fmt.Println(person)
	person.id = 20
	fmt.Println(person)
}
