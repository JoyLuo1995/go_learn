package main

import (
    "fmt"
)

func plusOne(digits []int) []int {
    n := len(digits)
    fmt.Println("length:", n)
    for i := n - 1; i >= 0; i-- {
        if digits[i] != 9 {
            digits[i]++            
            for j := i + 1; j < n; j++ {
                digits[j] = 0
            }
            return digits
        }
    }
    // digits 中所有的元素均为 9
    digits = make([]int, n+1)
    digits[0] = 1
    return digits
}

func main(){
    digits := []int{9, 9, 9}
    result := plusOne(digits)
    fmt.Println(result) // 输出: [1 0 0 0]
}