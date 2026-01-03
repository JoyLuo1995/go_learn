package main

import (
    "fmt"
)

func isValid(s string) bool {
    stack := []rune{}
    mapping := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    for _, ch := range s {
		//如果有左括号，先把所有左括号放到stack里
        if ch == '(' || ch == '[' || ch == '{' {
            stack = append(stack, ch)
        } else {
			//如果stack长度为0，说明没有左括号
			//如果stack的栈顶的左括号与当前ch键对应的值不相等，那就证明不匹配
			// ch只有可能是右括号，所以ch对应的值应该是左括号
            if len(stack) == 0 || stack[len(stack)-1] != mapping[ch] {
                return false
            }
			//如果长度不为零，且栈顶的左括号与当前ch键对应的值相等
			//那就证明匹配成功
			//此时栈顶的左括号已经判定过了，所以栈顶的左括号弹出

            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
    fmt.Println(isValid(")("))
    fmt.Println(isValid("()[]{}"))
    fmt.Println(isValid("(]"))
    fmt.Println(isValid("([)]"))
    fmt.Println(isValid("{[]}"))     
}