package main

import (
	"fmt"
	"sort"
)

type StudentScore struct { //定义一个结构体类型
	name  string
	score int
}

type StudentScores []StudentScore //定义一个StudentScore类型的结构体StudentScores

func (s StudentScores) Len() int { //获取元素数量
	return len(s)
}
func (s StudentScores) Less(i, j int) bool { //升序排序
	return s[i].score < s[j].score
}
func (s StudentScores) Swap(i, j int) { //交换位置
	s[i], s[j] = s[j], s[i]
}

func main() {
	//结构体的实例化
	x := StudentScores{
		{"赵周", 91},
		{"钱吴", 97},
		{"孙郑", 88},
		{"李王", 87},
	}
	fmt.Println("初始顺序：")
	for _, v := range x {
		fmt.Println(v.name, ":", v.score)
	}
	fmt.Println()

	sort.Sort(x)
	fmt.Println("排序顺序：")
	for _, v := range x {
		fmt.Println(v.name, ":", v.score)
	}
	fmt.Println()

}
