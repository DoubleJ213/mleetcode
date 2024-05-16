package fxxk

import (
	"fmt"
	"testing"
)

//
//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
//实现 LRUCache 类：
//
//LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//进阶：你是否可以在O(1) 时间复杂度内完成这两种操作？

type node146 struct {
	val  int
	next *node146
	pre  *node146
}

type LRUCache struct {
	capacity int
	size     int
	values   map[int]int
	nodes    map[int]*node146
	first    *node146
	last     *node146
}

func LruConstructor(capacity int) LRUCache {
	tmpNode := &node146{val: -1, next: nil, pre: nil}
	nodesMap := make(map[int]*node146)
	nodesMap[-1] = tmpNode
	value := make(map[int]int)
	value[-1] = -2
	return LRUCache{
		capacity: capacity,
		size:     1,
		values:   value,
		nodes:    nodesMap,
		first:    tmpNode,
		last:     tmpNode,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.values[key]; ok {
		thisNode := this.nodes[key]
		if this.last == thisNode {
			return v
		}
		if this.first == thisNode {
			this.first = this.first.next
		}
		thisNode.pre.next, thisNode.next.pre = thisNode.next, thisNode.pre
		this.last.next, thisNode.pre, thisNode.next = thisNode, this.last, nil
		this.last = thisNode
		return v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.values[key]
	this.values[key] = value
	if ok {
		//	已存在，更新
		thisNode := this.nodes[key]
		if this.last == thisNode {
			return
		}
		thisNode.pre.next, thisNode.next.pre = thisNode.next, thisNode.pre
		if this.first == thisNode {
			this.first = thisNode.next
		}
		thisNode.next, thisNode.pre, this.last = nil, this.last, thisNode
	} else {
		// 不存在，添加
		thisNode := &node146{val: value, pre: nil, next: nil}
		thisNode.pre, this.last.next = this.last, thisNode
		this.last = thisNode
		this.nodes[key] = thisNode
		if this.capacity > this.size {
			//	容量充足
			this.size += 1
		} else {
			//	容量不足
			delete(this.nodes, this.first.val)
			delete(this.values, this.first.val)
			this.first = this.first.next
		}
	}
}

func TestAl146(t *testing.T) {
	obj := LruConstructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(3))
}
