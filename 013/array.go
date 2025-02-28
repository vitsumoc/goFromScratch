package main

import "unsafe"

type Array struct {
	len int            // 当前元素数量
	cap int            // 底层数组容量
	p   unsafe.Pointer // 指向数组的指针
}
