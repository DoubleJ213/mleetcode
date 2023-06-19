package learn

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func f_1(a int) {
	a = 2
}

func f_1_1(a *int) {
	*a = 2
}

func f_2(s string) {
	s = "cba"
}

func f_2_1(s *string) {
	*s = "cba"
}

func f_3(v [2]string) {
	v[0] = "haha"
}

func f_3_s(v []string) {
	v[0] = "haha"
}

func f_3_1(v []string) {
	v = nil
}

func f_3_2(v *[]string) {
	*v = nil
}

func f_4(m map[int]int) {
	m[1] = 3
	m[3] = 1
}

func f_4_1(m map[int]int) {
	m = nil
}

func f_4_2(m *map[int]int) {
	*m = nil
}

func f_5(b []byte) {
	b[0] = 0b00101000
	//b[0] = 050
	//b[0] = 0x40
	var a int
	var c byte
	a = 050
	c = 050
	fmt.Printf("aaa %v\n", a)
	fmt.Printf("ccc %v\n", c)
	fmt.Printf("ccc string %s\n", string(c))
}

func f_5_1(b []byte) {
	b = bytes.Replace(b, []byte("1"), []byte("a"), -1)
}

func f_5_2(b *[]byte) {
	*b = bytes.Replace(*b, []byte("1"), []byte("a"), -1)
}

type why struct {
	s []string
}

func (ss why) SetV(s []string) {
	ss.s = s
}

func (ss *why) SetP(s []string) {
	ss.s = s
}

func (ss why) String() string {
	return strings.Join(ss.s, ",")
}

func TestZhi(*testing.T) {
	a := 1
	s := "abc"
	d := [...]string{"sd", "aa"}
	v := []string{"sd", "aa"}
	m := map[int]int{1: 1, 2: 2, 3: 3}
	f_1(a)
	f_2(s)
	f_3(d)
	f_3_s(v)
	f_4(m)
	fmt.Printf("%d,%s,%v,%v,%v\n", a, s, d, v, m)
	f_3_1(v)
	f_4_1(m)
	fmt.Printf("%d,%s,%v,%v\n", a, s, v, m)
	f_1_1(&a)
	f_2_1(&s)
	f_3_2(&v)
	f_4_2(&m)
	fmt.Printf("%d,%s,%v,%v\n", a, s, v, m)
	b := []byte("12145178")
	f_5(b)
	fmt.Printf("%s\n", b)
	f_5_1(b)
	fmt.Printf("%s\n", b)
	f_5_2(&b)
	fmt.Printf("%s\n", b)
	ss := &why{}
	ss.SetV([]string{"abc", "efg"})
	fmt.Println(ss)
	ss.SetP([]string{"abc", "efg"})
	fmt.Println(ss)

	/*
	   1,abc,[haha aa],map[1:3 2:2 3:1] slice和map值传递是可以修改原数据的，但基本数据类型不行
	   1,abc,[haha aa],map[1:3 2:2 3:1] 整体赋值不会修改原数据，值得注意的是map是无序的
	   2,cba,[],map[] 传递指针始终会修改原数据
	   @2145178 同上
	   @2145178 使用bytes.Replace实际上还是赋值，所以不会修改原数据
	   @2a45a78 使用指针传参就可以了
	            类的成员函数定义为传值的方式，不会修改原数据（原数据为空）
	   abc,efg  类的成员函数定义为传指针的方式，可以修改原成员变量

	   成功: 进程退出代码 0.
	*/

	/*
	   golang传值和传引用
	   这里不会解释关于指针的情况，如果读者对C语言或者C++的指针比较了解，那么就能更好地理解本文。

	   定义
	   对于代码
	   modify(a);
	   a.modify();
	   如果modify中对于a的修改不会改变传入的a的值，那么就是传值调用；否则，是传引用。
	   传值调用是将传入的变量在内存中复制一份进行操作，所以本质是存储在不同内存地址的不同变量。
	   传引用是将传入变量的内存地址，在函数操作中，通过内存地址将变量取出进行操作，所以本质是存储在同一个内存地址的相同变量
	*/

	as := "asdasdasda"
	for i, i2 := range as {
		//int8 rune
		fmt.Print(i2)
		fmt.Print(i)
	}

	fmt.Println(NodeID{0xFF}.PrefixLen())
	fmt.Println(NodeID{0x7F}.PrefixLen())
	fmt.Println(NodeID{0x3F}.PrefixLen())
	fmt.Println(NodeID{0x1F}.PrefixLen())
	fmt.Println(NodeID{0x0F}.PrefixLen())
	// ...
	fmt.Println(NodeID{0x00, 0x00, 0x05}.PrefixLen()) // 21
	// ...
	fmt.Println(NodeID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}.PrefixLen()) // 159
	fmt.Println(NodeID{}.PrefixLen())                                                           // 159

	//fmt.Println(PrefixLen("asdasdasda"))

}

const IdLength = 20

type NodeID [IdLength]byte

func (node NodeID) PrefixLen() (ret int) {
	for i := 0; i < IdLength; i++ {
		for j := 0; j < 8; j++ {
			if (node[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return IdLength*8 - 1
}

//func PrefixLen所做的是查找node开始的连续零的数量。外部循环每次执行一个字节，而内部循环处理不同的位:当j=0时，它向右移动7，因此从左边得到第一个位，当j=1时，它向右移动6，得到第二个位，以此类推。当遇到1的位时，它返回已经检查过的位的计数。

func TestL1(t *testing.T) {
	hehe := 'b' - 'a'
	//单引号，表示byte类型或rune类型，对应 uint8和int32类型，默认是 rune 类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。
	fmt.Printf("%v", hehe)

	/*
		hehe1 := "b" - "a"
		双引号，是字符串，实际上是字符数组。
		fmt.Printf("%v", hehe1)
		Golang限定字符或者字符串一共三种引号，单引号（’’)，双引号("") 以及反引号(``)。反引号就是标准键盘“Esc”按钮下面的那个键。
		对应的英文是：Single quote、Double quote、Back quote。
		单引号，表示byte类型或rune类型，对应 uint8和int32类型，默认是 rune 类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。
		双引号，才是字符串，实际上是字符数组。可以用索引号访问某字节，也可以用len()函数来获取字符串所占的字节长度。
		反引号，表示字符串字面量，但不支持任何转义序列。字面量 raw literal string 的意思是，你定义时写的啥样，它就啥样，你有换行，它就换行。你写转义字符，它也就展示转义字符。
	*/

}
