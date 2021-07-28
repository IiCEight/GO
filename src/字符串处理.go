package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

const N = 2e5 + 5

var s []byte

// var s []byte

// var s string

func main() {
	fmt.Scanf("%s", &s)
	n := len(s)
	// fmt.Println(string(s[0:n]))

	var last = 0
	for i := 0; i < n; i++ {
		if s[i] == ',' {
			t := s[last:i]
			var a []int
			cntn, cntu, cntl, cntchar := 0, 0, 0, 0
			for j := last; j < i; j++ {
				if unicode.IsDigit(rune(s[j])) {
					cntn++
					a = append(a, int(s[j]-'0'))
				} else if unicode.IsUpper(rune(s[j])) {
					cntu++
				} else {
					cntl++
				}
				if unicode.IsLetter(rune(s[j])) {
					cntchar++
				}

				if cntn == i-last { //方法一：判断字符串是否全是数字
					fmt.Println("Substring: ", string(t), " is all digits!")
					sort.Ints(a) //方法五：如果字符串全是数字组成，则按照数字大小升序排序并装入整型数组中。最后按升序打印出该数组
					fmt.Println("Let me sort it out:")
					for _, c := range a {
						fmt.Printf("%d ", c)
					}
					fmt.Printf("\n")
					// fmt.Println(a)
				} else if cntu == i-last { //方法二：判断字符串是否是大写字母
					fmt.Println("Substring: ", string(t), " is all uppercase letters!")
				} else if cntl == i-last { //方法三：判断字符串是否全是小写字母
					fmt.Println("Substring: ", string(t), " is all lowercase letters!")
				}

				if cntchar == i-last { //方法四：如果字符串全是字母，将其中所有小写字母转换为大写字母。
					fmt.Println("This substring ", string(t), " is all letter, Let me change it all into uppercase:")
					fmt.Println(strings.ToUpper(string(t)))
				}
			}
			last = i + 1

			fmt.Printf("\n")
		}
	}

	fmt.Println("Thank You!!!")
}
