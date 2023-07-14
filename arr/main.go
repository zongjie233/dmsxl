package main

import "fmt"

func main() {
	slice := []int{1, 3, 3, 6, 7}
	nums := []int{-7, -3, 2, 3, 11}
	numsFor4 := []int{2, 3, 1, 2, 4, 3}
	numsfor4 := []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(erfenchazhao(slice, 7)) //4
	// fmt.Println(yichuyuansu(slice, 3))  //3
	// fmt.Println(demo3(nums))
	fmt.Println(demo4(numsFor4, 7))  // 2
	fmt.Println(demo4(numsfor4, 11)) // 0
	fmt.Println(luoxuan(3))
}
