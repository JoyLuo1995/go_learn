package main

import (
    "fmt"
)

func LongestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    prefix := strs[0]
    for _, s := range strs[1:] {
        i := 0
        for i < len(prefix) && i < len(s) && prefix[i] == s[i] {
            i++
        }
        prefix = prefix[:i]
        if prefix == "" {
            return ""
        }
    }
    return prefix
}

func main() {
    strs := []string{"flower", "flow", "flight"}
    prefix := LongestCommonPrefix(strs)
    fmt.Println(prefix) // Output: "fl"
}