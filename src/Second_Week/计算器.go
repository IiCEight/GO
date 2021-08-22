//原理：边中缀表达式转换成后缀表达式边进行计算
//时间复杂度O(N)
//写得有点丑，dalao别嫌弃呀
package main

import (
	"fmt"
	"unicode"
)

const N int = 2e2 + 5

var n int

type num struct {
	st  [N]float64
	top int
}

type op struct {
	st  [N]byte
	top int
}

func (op *op) calc(num *num) { //计算部分
	switch op.st[op.top] {
	case '+':
		num.st[num.top-1] = num.st[num.top-1] + num.st[num.top]
		num.top--
	case '-':
		num.st[num.top-1] = num.st[num.top-1] - num.st[num.top]
		num.top--
	case '*':
		num.st[num.top-1] = num.st[num.top-1] * num.st[num.top]
		num.top--
	case '/':
		num.st[num.top-1] = num.st[num.top-1] / num.st[num.top]
		num.top--
	}
}

func main() {
	for {
		num := num{top: 0}
		op := op{top: 0}
		fmt.Println("Give me the correct expression in one line\nI will calculate the result(alt + F4 to quit)")
		var s string
		fmt.Scanln(&s)
		n = len(s)

		for i := 0; i < n; i++ { //遍历一遍目标表达式
			var res int = 0
			if unicode.IsDigit(rune(s[i])) {
				for i < n && unicode.IsDigit(rune(s[i])) {
					res = res*10 + int(s[i]-'0')
					i++
				}
				num.top++
				num.st[num.top] = float64(res)
				i--
			} else if s[i] == '(' {
				op.top++
				op.st[op.top] = '('
			} else if s[i] == '+' {
				if op.st[op.top] == '*' {
					num.st[num.top-1] = num.st[num.top-1] * num.st[num.top]
					num.top--
					op.top--
				} else if op.st[op.top] == '/' {
					num.st[num.top-1] = num.st[num.top-1] / num.st[num.top]
					op.top--
					num.top--
				}
				op.top++
				op.st[op.top] = '+'
			} else if s[i] == '-' {
				if op.st[op.top] == '*' {
					num.st[num.top-1] = num.st[num.top-1] * num.st[num.top]
					num.top--
					op.top--
				} else if op.st[op.top] == '/' {
					num.st[num.top-1] = num.st[num.top-1] / num.st[num.top]
					op.top--
					num.top--
				}
				op.top++
				op.st[op.top] = '-'
			} else if s[i] == '*' {
				op.top++
				op.st[op.top] = '*'
			} else if s[i] == '/' {
				op.top++
				op.st[op.top] = '/'
			} else if s[i] == ')' {
				for op.st[op.top] != '(' {
					op.calc(&num)
					op.top--
				}
				op.top--
			}
		}

		for op.top > 0 { // 处理符号栈中残留的符号
			op.calc(&num)
			op.top--
		}

		fmt.Println("Result is ", num.st[num.top])
	}
}
