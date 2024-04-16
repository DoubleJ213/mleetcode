package learn

import (
	"fmt"
	"testing"
	"unsafe"
)

type Coder interface {
	code()
}

type Gopher struct {
	language string
	name     string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func Test_St1(t *testing.T) {
	/*	var c coder = &Gopher{"Go"}
		c.code()
		c.debug()

		//var b coder = Gopher{"Golang b"}
		// 会说Gopher 没有实现coder
		//b.code()
		//b.debug()
		//_= b

		d := Gopher{language: "Golang d"}
		d.code()
		d.debug()

		e := &Gopher{language: "Golang e"}
		e.code()
		e.debug()*/

	/*var f Coder
	fmt.Println(f == nil)
	fmt.Printf("f: %T, %v\n", f, f)

	var g *Gopher
	fmt.Println(g == nil)

	f = g
	fmt.Println(f == nil)
	fmt.Printf("f: %T, %v\n", f, f)*/

	/*	var h Coder
		fmt.Println(&h == nil)
		fmt.Printf("&h: %T, %v\n", &h, &h)

		var g Gopher
		fmt.Println(&g == nil)
		g.language = "1"
		fmt.Printf("&g: %T, %v\n", &g, &g)

		var i *Gopher
		fmt.Println(i == nil)
		fmt.Printf("i: %T, %v\n", i, i)
		i = new(Gopher) // 内存初始化过，有零值 不为nil
		fmt.Println(i == nil)
		fmt.Println(i.language)
		fmt.Println(i.name)
		i.language = "2"
		//invalid memory address or nil pointer dereference
		fmt.Printf("i: %T, %v\n", i, i)
	*/
	/*	var a int = 100
		var p *int = &a
		//var 指针变量名称 *指针类型
		// 通过指针修改变量的值
		*p = 666
		//var p *int = &a  =>  p = &a  => *p = *&a  => *p = a => p取值后变成int类型 => p 对应为int类型的指针
		fmt.Printf("a变量的值:%v\ta变量的内存地址:%v\n", a, p)
		//a变量的值:666	a变量的内存地址:0xc000018098
		//大多数指针类型会写成*T，表示是“一个指向T类型变量的指针”。

		fmt.Println("十六进制 unsafe.Pointer(p): ", unsafe.Pointer(p))
		fmt.Println("十进制 uintptr(unsafe.Pointer(p)):  ", uintptr(unsafe.Pointer(p)))
		fmt.Println("unsafe.Pointer(uintptr(unsafe.Pointer(p))):  ", unsafe.Pointer(uintptr(unsafe.Pointer(p))))
	*/

	/*	length := 6
		arr := make([]int, length)
		for i := 0; i < length; i++ {
			arr[i] = i
		}
		fmt.Println(arr)
		// [0 1 2 3 4 5]
		// 取slice的第5个元素：通过计算第1个元素 + 4 个元素的size 得出
		fmt.Println("unsafe.Sizeof(arr[0]): ", unsafe.Sizeof(arr[0]))
		//uintptr 存的是 10进制地址
		fmt.Println("uintptr(unsafe.Pointer(&arr[0])): ", uintptr(unsafe.Pointer(&arr[0])))
		end := unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + 4*unsafe.Sizeof(arr[0]))

		fmt.Println(*(*int)(end)) // 4
		fmt.Println(arr[4])       // 4
		// 和 pb := &x.b 等价
	*/

	//fmt.Println("unsafe.Offsetof(x.b): ", unsafe.Offsetof(x.b))
	//fmt.Println("unsafe.Sizeof(x.c): ", unsafe.Sizeof(x.c))
	/*
	   int, uint, uintptr	1个机器字
	   已知 int 64位 类似 int64  64/8  8个字节
	   如果是 32位 那就是 int32 32/8  4个字节
	   []T	3个机器字（data、len、cap） 64 那就是8*3
	   换个想法 1个机器字 其实就是1个指针地址存的容量
	*/

	//fmt.Println(unsafe.Pointer(&x))
	//fmt.Println(uintptr(unsafe.Pointer(&x)))
	//a2 := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	//fmt.Println("a2", a2)
	////a3 := fmt.Sprintf("%x", a2) //string 再转 int
	////fmt.Println("a3", a3)
	//a3 := strconv.FormatInt(int64(a2), 16)
	////func FormatInt(i int64, base int) string {}
	//fmt.Println("a3", a3)
	//a4, _ := strconv.ParseInt(a3, 16, 64)
	//fmt.Println("a4", a4)
	//fmt.Println(unsafe.Pointer(uintptr(a3)))

	//10958338
	/*	fmt.Println("--1--", &x)
		fmt.Println("--2--", unsafe.Pointer(&x))
		pb := (*int16)(unsafe.Pointer(
			uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

		//pb1 := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x))))

		*pb = 42
		fmt.Println(x.b) // "42"*/

	//var hexNumber int = 0x1F
	//fmt.Println(hexNumber)
	//虽然能定义 这个十六进制数 其实存储时会被转换成十进制的31，并赋值给hexNumber变量。

	//通常我们的操作系统为 64 位，也就是说操作系统操作内存都是以 8 字节为基础进行操作。
	//我们定义了一个大小为 4 字节的字段实，如果我们以 4 作为偏移量去操作后续的字段，
	// 会对操作系统管理内存带来诸多不便，所以 go 会将剩余的 4 字节填充，以达到对齐内存的目的。
	//fmt.Println(unsafe.Sizeof(xx.aaaa)) // 1
	//fmt.Println(unsafe.Offsetof(xx.a)) // 2
	//fmt.Println(unsafe.Sizeof(xx.a)) // 2
	//fmt.Println(unsafe.Offsetof(xx.b))  // 8

	//fmt.Println(unsafe.Sizeof(yy.aaaa)) // 1
	//fmt.Println(unsafe.Offsetof(yy.a))  // 2
	//fmt.Println(unsafe.Sizeof(yy.a))    // 2
	//fmt.Println(unsafe.Offsetof(yy.b))  // 8

	/*	fmt.Println(unsafe.Sizeof(zz.aaaa))  // 1
		fmt.Println(unsafe.Offsetof(zz.aaa)) // 1
		fmt.Println(unsafe.Sizeof(zz.aaa))   // 1
		fmt.Println(unsafe.Offsetof(zz.a))   // 2
		fmt.Println(unsafe.Sizeof(zz.a))     // 2
		fmt.Println(unsafe.Offsetof(zz.aa))  // 4
		fmt.Println(unsafe.Sizeof(zz.aa))    // 1
		fmt.Println(unsafe.Offsetof(zz.b))   // 8*/

	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)

	//现在我们通过指针运算给b变量赋值为10
	fmt.Println("--", unsafe.Offsetof(w.b))
	//b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + 2*unsafe.Sizeof(w.b))
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))
	*((*int)(b)) = 10
	//此时结果就变成了10，0
	fmt.Println(w.b, w.c)
}

type W struct {
	b int32
	c int64
}

var zz struct {
	aaaa bool  //1个字节
	aaa  bool  //1个字节
	a    int16 //2 字符
	aa   bool  //1个字节
	b    string
}

var yy struct {
	aaaa bool  //1个字节
	a    int64 //64位系统  int int64
	b    string
}

var xx struct {
	aaaa bool //1个字节
	a    int16
	//aa  int8
	// 8/8 (intN  N/8)个字节
	//aaa int8
	b string
}

/*
bool	1个字节
intN, uintN, floatN, complexN	N/8个字节（例如float64是8个字节）
int, uint, uintptr	1个机器字
*T	1个机器字
string	2个机器字（data、len）
[]T	3个机器字（data、len、cap）
map	1个机器字
func	1个机器字
chan	1个机器字
interface	2个机器字（type、value）
*/

var x struct {
	a  bool
	aa int8
	b  int16
	c  []int
}

func demo(v1, v2 *int) int { //传入的参数是int类型指针
	*v1 = 5
	*v2 = 6
	return *v1 * *v2
}
