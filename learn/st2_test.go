package learn

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type myWriter struct {
}

//Write(p []byte) (n int, err error)

//func (*myWriter) Write(p []byte) (n int, err error) {
func (myWriter) Write(p []byte) (n int, err error) {
	fmt.Println("heheda")
	return 0, nil
}

/*
type Person struct {
	age int
}
func (p Person) howOld() int {
	return p.age
}


func (p *Person) growUp() {
	p.age += 1
}*/

type Person interface {
	growUp()
}
type Student struct {
	age  int
	name string
}

////String() string
func (s *Student) String() string {
	return fmt.Sprintf("My Name Is %s", s.name)
}

func (p Student) growUp() {
	p.age += 1
	return
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter uintptr
	_type uintptr
	link  uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

func Test_St2(t *testing.T) {
	/*	// qcrao 是值类型
		qcrao := Person{age: 18}

		// 值类型 调用接收者也是值类型的方法
		fmt.Println(qcrao.howOld())

		// 值类型 调用接收者是指针类型的方法
		qcrao.growUp()
		fmt.Println(qcrao.howOld())

		// ----------------------

		// stefno 是指针类型
		stefno := &Person{age: 100}

		// 指针类型 调用接收者是值类型的方法
		fmt.Println(stefno.howOld())

		// 指针类型 调用接收者也是指针类型的方法
		stefno.growUp()
		fmt.Println(stefno.howOld())*/

	/*// 检查 *myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = (*myWriter)(nil)

	// 检查 myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = myWriter{}
	*/
	/*	var qcrao = Person(Student{age: 18})

		iface := (*iface)(unsafe.Pointer(&qcrao))

		fmt.Printf("iface.tab.hash = %#x\n", iface.tab.hash)
		fmt.Println(iface.data)
		fmt.Println(*(*int)(iface.data))

		var i interface{} = new(Student)
		//func new(Type) *Type
		//new(Student) 创建一个 Student 类型的指针
		s := i.(*Student)
		fmt.Println(s)

		num := 12

		switch {
		case num%2 == 0 && num%3 == 0:
			fmt.Println("The number is divisible by both 2 and 3.")
			fallthrough
		case num%2 == 0:
			fmt.Println("The number is divisible by 2.")
		case num%3 == 0:
			fmt.Println("The number is divisible by 3.")
		default:
			fmt.Println("The number is not divisible by either 2 or 3.")
		}

		aa := Student{age: 18, name: "hello"}
		fmt.Println(aa)
	*/
	ch := make(chan int)
	go goroutineA(ch)
	go goroutineB(ch)
	//go goroutineC(ch)
	go goroutineD(ch)
	ch <- 3
	time.Sleep(time.Second)
}

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("G1 received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <-b
	fmt.Println("G2 received data: ", val)
	return
}

func goroutineC(b <-chan int) {
	val := <-b
	fmt.Println("G3 received data: ", val)
	return
}
func goroutineD(b <-chan int) {
	val := <-b
	fmt.Println("G4 received data: ", val)
	return
}
