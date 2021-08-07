package main

import (
	"fmt"
)

const N = 2e5 + 5

type node struct {
	val int
	f   *node
	r   *node
}

var pre *node = nil
var iterator *node = nil //迭代器iterator

func add() { // 增
	p := new(node)
	fmt.Println("which value do you want to add?")
	fmt.Scanln(&p.val)
	p.f = pre
	p.r = iterator
	pre.r = p
	pre = p
}

func find() *node { // 查
	var which int
	fmt.Scanln(&which)
	cnt := 0
	ans := false
	i := iterator.r
	for ; i != iterator; i = i.r {
		// fmt.Printf("%d\n", i.val)
		cnt++
		if i.val == which {
			ans = true
			break
		}
	}
	if ans {
		fmt.Printf("It's in %d node\n", cnt)
		return i
	} else {
		fmt.Println("No found!")
		return nil
	}
}

func Delete() { //删
	fmt.Println("Which value do you want to delete?")
	p := find()
	if p != nil {
		fmt.Println("Delete Successfully!")
		p.f.r = p.r
		p.r.f = p.f
		p = nil
	}
}

func modify() { // 改
	fmt.Println("Which value do you want to modify?")

	p := find()
	if p != nil {
		fmt.Println("Which value do you want it to be?")
		var value int
		fmt.Scanln(&value)
		fmt.Println("Modify Successfully!")
		p.val = value
	}
}

func main() {
	fmt.Println("Welcome!!!")
	p := new(node)
	p.f = p
	p.r = p
	iterator = p
	pre = p
	for {
		fmt.Println("1.add 2.delete 3.modify 4.find Else.quit")
		var choice int
		fmt.Scanln(&choice)
		// fmt.Print(choice)
		// fmt.Scanlnf("%d", &choice)
		// fmt.Println("hhhh")
		flag := false
		switch choice {
		case 1:
			add()
		case 2:
			Delete()
		case 3:
			modify()
		case 4:
			fmt.Println("Which value do you want to find?")
			find()
		default:
			flag = true
		}
		if flag {
			break
		}
	}

}
