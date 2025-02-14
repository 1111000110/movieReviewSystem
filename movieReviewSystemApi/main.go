package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	rotate(x, 3)
	fmt.Println(x)
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	nums = append(nums[len(nums)-k:len(nums)], nums[:len(nums)-k]...)
	fmt.Println(nums)
}
