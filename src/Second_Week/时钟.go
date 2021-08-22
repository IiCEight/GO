package main

import (
	"fmt"
	"time"
)

type mytime struct {
	day, month, sec, minu, hour, year int
}

var bigmonth = map[int]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

var t mytime

func (t *mytime) flow() { //时间进位写得有点繁琐了，希望dalao别嫌弃
	t.sec++
	if t.sec >= 60 {
		t.sec = 0
		t.minu++
		if t.minu >= 60 {
			t.minu = 0
			t.hour++
			if t.hour >= 24 {
				t.hour = 0
				t.day++
				if t.month == 2 { //2月特判
					if (0 == t.year%4 && 0 != t.year%100) || (0 == t.year%400) { //润年
						if t.day > 29 {
							t.day = 1
							t.month++
							if t.month >= 12 {
								t.month = 1
								t.year++
							}
						}
					} else {
						if t.day > 28 {
							t.day = 0
							t.month++
							if t.month >= 12 {
								t.month = 1
								t.year++
							}
						}
					}
				} else {
					if t.day > bigmonth[t.month] {
						t.day = 1
						t.month++
						if t.month >= 12 {
							t.month = 1
							t.year++
						}
					}
				}
			}
		}
	}
}

func cur() {
	fmt.Println("Input Year")
	fmt.Scanln(&t.year)
	fmt.Println("Input Month")
	fmt.Scanln(&t.month)
	fmt.Println("Input Day")
	fmt.Scanln(&t.day)
	fmt.Println("Input hour")
	fmt.Scanln(&t.hour)
	fmt.Println("Input minuue")
	fmt.Scanln(&t.minu)
	fmt.Println("Input second")
	fmt.Scanln(&t.sec)
	for {
		fmt.Printf("%d-%d-%d %02d:%02d:%02d\n", t.year, t.month, t.day, t.hour, t.minu, t.sec)
		t.flow()
		time.Sleep(time.Second)
	}
}

func now() {
	for {
		t := time.Now()
		fmt.Printf("%d-%d-%d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())

		time.Sleep(time.Second)
	}
}

func main() {
	flag := false
	for {
		fmt.Println("1.Input time 2.Now Else.quit")
		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			cur()
		} else if choice == 2 {
			now()
		} else {
			flag = true
		}
		if flag {
			break
		}
	}
}
