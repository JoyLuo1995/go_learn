package main

import (
	"fmt"
)

func Remove_Duplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    i := 0
    for j := 1; j < len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }
    return i + 1
}

func main() {
    nums := []int{1, 1, 2, 2, 3, 3, 3, 4}
    k := Remove_Duplicates(nums)
    fmt.Println("k =", k)
    fmt.Println("前 k 个元素:", nums[:k])  // 输出前 k 个元素
}