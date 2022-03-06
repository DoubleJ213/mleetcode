package main

import (
	"fmt"
	"runtime"
	"testing"
)

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func paseStudent() map[string]*student {
	m := make(map[string]*student)
	stus := [...]student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range stus {
		//别想太复杂，for遍历时，变量stu指针不变，同一块内存，来回赋值而已。
		//每次遍历仅进行struct值拷贝，故m[stu.Name]=&stu实际上一致指向同一个指针，
		//最终该指针的值为遍历的最后一个struct的值拷贝
		m[stu.Name] = &stu
		fmt.Print(i)

		m[stu.Name] = &stus[i]
	}
	return m
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestA1(t *testing.T) {
	//defer_call()

	//fmt.Print(3 | 4)
	//fmt.Print(0b011)
	//fmt.Print(0b100)
	//fmt.Print(0b111)
	//
	//numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	//maxIndex2 := len(numbers2) - 1
	//for i, e := range numbers2 {
	//	if i == maxIndex2 {
	//		numbers2[0] += e
	//	} else {
	//		numbers2[i+1] += e
	//	}
	//}
	//fmt.Println(numbers2)
	//
	//numbers3 := []int{1, 2, 3, 4, 5, 6}
	//maxIndex3 := len(numbers3) - 1
	//for i, e := range numbers3 {
	//	if i == maxIndex3 {
	//		numbers3[0] += e
	//	} else {
	//		numbers3[i+1] += e
	//	}
	//}
	//fmt.Println(numbers3)
	//range拷贝的是切片结构体的副本，切片结构体中有指向底层数组的指针。
	//这个副本与原本的指针都指向同一个底层数组。

	//students := paseStudent()
	//for k, v := range students {
	//	fmt.Printf("key=%s,value=%v \n", k, v)
	//}

	//numbers1 := []int{1, 2, 3, 4, 5, 6}
	//for i := range numbers1 {
	//	if i == 3 {
	//		//numbers1[i] |= i
	//		numbers1[i] = numbers1[i] | i
	//	}
	//}
	//fmt.Println(numbers1)
	//value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//value1 := [...]int{0, 1, 2, 3, 4, 5, 6}
	////如果switch表达式的结果值是无类型的常量，
	////比如1 + 3的求值结果就是无类型的常量4，
	////那么这个常量会被自动地转换为此种常量的默认类型的值，
	////比如整数4的默认类型是int，又比如浮点数3.14的默认类型是float64。
	////因此，由于上述代码中的switch表达式的结果类型是int，
	////而那些case表达式中子表达式的结果类型却是int8，它们的类型并不相同，
	////所以这条switch语句是无法通过编译的。
	//switch 1 + 2 {
	//case value1[0], value1[1]:
	//	fmt.Println("0 or 1")
	//case value1[2], value1[3]:
	//	fmt.Println("2 or 3")
	//	fallthrough
	//	//go语言fallthrough的用法心得：
	//	//Go里面 switch 默认相当于每个case最后带有break，
	//	//匹配成功后不会自动向下执行其他case，而是跳出整个switch,
	//	//但是可以使用fallthrough强制执行后面的case代码。
	//	//fallthrough不能用在switch的最后一个分支 fallthrough到下一个case块时，
	//	//不执行case匹配检查！不执行case匹配检查！不执行case匹配检查！
	//case value1[4], value1[5], value1[6]:
	//	fmt.Println("4 or 5 or 6")
	//}
	//
	//value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch value2[4] {
	//case 0, 1:
	////case 语句的表达式是无类型的常量，自动转换为switch 表达式的结果类型
	////当然了，如果这里说的自动转换没能成功，那么switch语句照样通不过编译
	//	fmt.Println("0 or 1")
	//case 2, 3:
	//	fmt.Println("2 or 3")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	//runtime.GOMAXPROCS(1)
	//wg := sync.WaitGroup{}
	//wg.Add(20)
	//for i := 0; i < 10; i++ {
	//	time.Sleep(1000)
	//	go func() {
	//		fmt.Println("i1: ", i)
	//		//这里没传参，爱是啥是啥，而且是i一直加，加到10退出循环
	//		wg.Done()
	//	}()
	//}
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		fmt.Println("i2: ", i)
	//		//这里是把i传进来了，不管怎么打印，一定是0到9都有
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	//
	//fmt.Print("done")

	teacher := Teacher{}
	teacher.ShowB()

	//golang中没有继承，只有组合
	teacher.ShowA()
	//这里People是匿名组合People。
	//被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法，
	//但他们的方法(ShowA())调用时接受者并没有发生变化。
	//这里仍然是People。毕竟这个People类型并不知道自己会被什么类型组合，
	//当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。
	//因此这里执行t.ShowA()时，在执行ShowB()时该函数的接受者是People，而非Teacher

	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
	//
	//runtime.GOMAXPROCS(1)
	//int_chan := make(chan int)
	////string_chan := make(chan string)
	//int_chan <- 1  得先把取管道怼好才能写吗，不是，但是写完不取阻塞，取了没写也阻塞
	//go func() {
	//	int_chan <- 1
	//}()
	//
	//select {
	//case value := <-int_chan:
	//	fmt.Println(value)
	//	//case value := <-string_chan:
	//	//	panic(value)
	//	//	fmt.Println(value)
	//}
	fmt.Print("done")
	//单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：
	//select 中只要有一个case能return，则立刻执行。
	//	当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
	//	如果没有一个case能return则可以执行”default”块。
	//	此考题中的两个case中的两个chan均能return，则会随机执行某个case块。故在执行程序时，有可能执行第二个case，触发异常
}
