package main

import "fmt"

func main() {
	v := twoSum([]int{1, 2, 3}, 5)
	fmt.Println(v)
}

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, v := range nums {
		if _, ok := hm[v]; ok {
			return []int{i, hm[v]}
		}
		hm[target-v] = i
	}
	return nil
}
