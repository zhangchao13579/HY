package main

//输入一个字符串，判断是否为回文串，如果是，则返回true；如果不是，则返回false.

import "fmt"

func main() {
	var a string
	fmt.Print("输入你的字符串：")
	_, err := fmt.Scanln(&a)
	if err != nil {
		return
	}

	var b = []rune(a)
	var l = len(b)
	for i := 0; i < l/2; i++ {
		c := b[i]
		b[i] = b[l-i-1]
		b[l-i-1] = c
	}

	d := string(b)

	if d == a {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}

}
