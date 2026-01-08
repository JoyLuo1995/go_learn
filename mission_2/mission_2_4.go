package main

import (
	"fmt"
	"time"
)


func main() {

	// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
	// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	// 考察点 ：通道的基本使用、协程间通信。
	ch := make(chan int)

	// 生产者协程
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond) // 模拟一些工作
		}
		close(ch) // 关闭通道，表示不再发送数据
	}()
	// 消费者协程
	go func() {
		for num := range ch {
			fmt.Println("Received:", num)
		}
	}()

	time.Sleep(2 * time.Second)


	// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	// 考察点 ：通道的缓冲机制。
	bufferedCh := make(chan int, 10) // 创建一个缓冲区大小为10的通道

	// 生产者协程
	go func() {
		for i := 1; i <= 100; i++ {
			bufferedCh <- i
			fmt.Println("Sent:", i)
		}
		close(bufferedCh) // 关闭通道，表示不再发送数据
	}()

	// 消费者协程
	go func() {
		for num := range bufferedCh {
			fmt.Println("Received:", num)
		}
	}()

	time.Sleep(5 * time.Second)

}