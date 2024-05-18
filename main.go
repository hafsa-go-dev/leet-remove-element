package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{
		3, 2, 2, 3,
	}

	fmt.Println(removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {
	n := nums
	for {
		index := slices.Index(n, val)
		if index == -1 {
			break
		} else {
			n = slices.Delete(n, index, index+1)
		}
	}

	return len(n)

}
