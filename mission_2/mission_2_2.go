

package main

import (
    "fmt"
	"sync"
	"time"
)


// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。

func PrintOdd() {
    for i := 1; i <= 10; i += 2 {
		time.Sleep(100 * time.Millisecond)
        fmt.Println("Odd:", i)
    }
}

func PrintEven() {
    for i := 2; i <= 10; i += 2 {
		time.Sleep(100 * time.Millisecond)
        fmt.Println("Even:", i)
    }
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

func Task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	duration := time.Since(start)
	fmt.Printf("Task %d completed in %v\n", id, duration)
}


func main() {

	go PrintOdd()
	go PrintEven()
	time.Sleep(2 * time.Second)


	var wg sync.WaitGroup
	tasks := 5
	for i := 1; i <= tasks; i++ {
		wg.Add(1)
		go Task(i, &wg)
	}
	wg.Wait()
}

