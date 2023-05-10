package main

import "fmt"

func main() {
	slice := []int{1, 3, 3, 6, 7}
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(erfenchazhao(slice, 7)) //4
	fmt.Println(yichuyuansu(slice, 3))  //3
	fmt.Println(demo(nums))
}
