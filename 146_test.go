package mleetcode

import (
	"fmt"
	"testing"
)

//146. LRU 缓存机制

//LRUCache(int capacity)以正整数作为容量capacity初始化LRU缓存
//int get(int key)如果关键字key存在于缓存中，则返回关键字的值，否则返回-1 。
//void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
//当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

// 用一个双向链表来维护当前的数据，越靠后表示越后访问
type LinkNode struct {
	Next *LinkNode
	Prev *LinkNode
	Val  int
}

type LRUCache struct {
	start *LinkNode
	end   *LinkNode
	//当前长度
	size int
	//容量
	capacity int
	//key对应的value  key加到双端链表的val中
	values map[int]int
	//存储key对应的当前node节点
	nodes map[int]*LinkNode
}

func LruConstructor(capacity int) LRUCache {
	node := &LinkNode{Next: nil, Prev: nil, Val: -1}
	var lCache = LRUCache{
		start:    node,
		end:      node,
		size:     1,
		capacity: capacity,
		values:   make(map[int]int),
		nodes:    make(map[int]*LinkNode),
	}
	return lCache
}

func (this *LRUCache) Get(key int) int {
	//更新缓存
	elem, ok := this.values[key]
	if ok {
		tmp := this.nodes[key]
		if tmp.Next != nil {
			// 把当前节点的上一个节点和下一个节点连接起来
			tmp.Prev.Next, tmp.Next.Prev = tmp.Next, tmp.Prev
			if this.start == tmp {
				//如果当前节点为最前面的节点，更新start为当前的后一个
				this.start = tmp.Next
			}
			// 处理尾节点,双端链表，双端都需要处理
			this.end.Next, tmp.Prev = tmp, this.end
			tmp.Next = nil
			this.end = tmp
		}
		return elem
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.values[key]
	this.values[key] = value
	//node := &LinkNode{nil, nil, key}
	if ok {
		//	更新，并把这个值挪到最后
		tmp := this.nodes[key]
		if this.end == tmp {
			return
		}
		tmp.Prev.Next, tmp.Next.Prev = tmp.Next, tmp.Prev
		if this.start == tmp {
			this.start = tmp.Next
		}
		this.end.Next, tmp.Prev = tmp, this.end
		tmp.Next = nil
		this.end = tmp
	} else {
		//	没找到key,加到最后
		//  如果长度超了，把最前面的挤走
		//	如果长度不超，直接添加
		tmp := &LinkNode{Next: nil, Prev: nil, Val: key}
		this.end.Next, tmp.Prev, this.end = tmp, this.end, tmp
		this.nodes[key] = tmp
		if this.capacity > this.size {
			//	添加
			this.size += 1
		} else {
			//	添加并删除前一个
			delete(this.values, this.start.Val)
			delete(this.nodes, this.start.Val)
			this.start = this.start.Next
		}
		//this.end.Next, tmp.Prev, this.end = tmp, this.end, tmp
		//this.nodes[key] = tmp
	}
}

func TestAl146(t *testing.T) {
	//两个测试用例
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	//["LRUCache","put","get","put","get","get"]
	//[[1],[2,1],[2],[3,2],[2],[3]]

	/*obj := LruConstructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
	*/

	/*obj1 := LruConstructor(1)
	obj1.Put(2, 1)
	fmt.Println(obj1.Get(2))
	obj1.Put(3, 2)
	fmt.Println(obj1.Get(2))
	fmt.Println(obj1.Get(3))*/

	//["LRUCache","put","put","put","put","get","get"]
	//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]

	/*obj2 := LruConstructor(2)
	obj2.Put(2, 1)
	obj2.Put(1, 1)
	obj2.Put(2, 3)
	obj2.Put(4, 1)
	fmt.Println(obj2.Get(1))
	fmt.Println(obj2.Get(2))
	*/

	fmt.Println("aa")
}
