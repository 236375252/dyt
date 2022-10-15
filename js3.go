package main

import "fmt"

func main() {
	var n int
	fmt.Print("n=")
	fmt.Scanln(&n)
	len := n
	var nums []int = make([]int, len)
	for i := 0; i < n; i++ {
		fmt.Print("nums=")
		fmt.Scanln(&nums[i])
	}
	//输入囚号
	var x []int = make([]int, len)
	var x1 [10]int = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < n; i++ {
		for a := 0; a < 10; a++ {
			for b := 0; b < 10; b++ {
				for c := 0; c < 10; c++ {
					for d := 0; d < 10; d++ {
						if nums[i] == x1[a]*1000+x1[b]*100+x1[c]*10+x1[d] {
							x[i] = x1[a] + x1[b] + x1[c] + x1[d]
						}
					}
				}
			}
		}
	}
	//求对应的牢房号
	var m int
	for j := 0; j < n; j++ {
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				m = x[i-1]
				x[i-1] = x[i]
				x[i] = m
			}
		}
	}
	//调为增序
	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			if x[j] == x[k] {
				x[k] = 0
			}
		}
	}
	var l int
	for j := 0; j < n; j++ {
		for i := 1; i < n; i++ {
			if x[i] < x[i-1] {
				l = x[i-1]
				x[i-1] = x[i]
				x[i] = l
			}
		}
	}
	for i := 0; i < n; i++ {
		if x[i] == 0 {
			n--
		}
	}
	var x2 []int = make([]int, n)
	for i := 0; i < n; i++ {
		x2[i] = x[n+i]
	}
	//去除相同牢房号
	fmt.Println(n)
	fmt.Println(x2)
}
