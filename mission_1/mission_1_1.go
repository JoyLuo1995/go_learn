//给定一个非空整数数组，除了某个元素只出现一次以外，
// 其余每个元素均出现两次。找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

package main

import (
    "fmt"
)

func singleNumber(nums []int) []int {
    count := make(map[int]int)
    var x []int
    for _, num := range nums {
        count[num] = count[num] + 1
    }
    for num, c := range count {
        if c == 1 {
            x = append(x, num)
        }
    }
    if len(x) > 0 {
        return x
    } else {
        fmt.Println("No single number found")
        return nil
    }
}

func main() {
    nums := []int{2, 2, 1, 3, 4}
    fmt.Println(singleNumber(nums))
}