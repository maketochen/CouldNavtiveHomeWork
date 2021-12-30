package main

import (
	"fmt"
	"time"
)

func main() {
	/*编写一个小程序：
		给定一个字符串数组
	["I","am","stupid","and","weak"]
		用 for 循环遍历该数组并修改为
	["I","am","smart","and","strong"]*/

	fmt.Println("hello world")
	b := [...]string{"I", "am", "stupid", "and", "weak"}
	for i := 0; i < len(b); i++ {
		if i == 2 {
			b[i] = "smart"
		}
	}
	for index, value := range b {
		if value == "weak" {
			b[index] = "strong"
		}
	}
	fmt.Println(b)
	/*基于 Channel 编写一个简单的单线程生产者消费者模型：

	队列：
	队列长度 10，队列元素类型为 int
	生产者：
	每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
	消费者：
	每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞*/
	ch := make(chan int, 6)
	go product(ch)
	consumer(ch)

}

func product(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
		fmt.Printf("produce data:%d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("get data %d\n", i)
	}
}
