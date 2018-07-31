package main

import (
	"fmt"
)

type My struct {
	num int
}

func (self My) addOne() {
	self.num++
}

func (self *My) addTwo() {
	self.num += 2
}

func main() {
	my1 := My{1}
	my1.addOne()
	fmt.Println(my1.num)
	my1.addTwo()
	fmt.Println(my1.num)
}
