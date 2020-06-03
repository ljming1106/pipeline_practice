package pipeline

import "fmt"

/*
1、遍历切片（生产者）
2、计算平方值。
3、打印结果（消费者）
*/

// 流水线的特点
// 每个阶段把数据通过channel传递给下一个阶段。
// 每个阶段要创建1个goroutine和1个通道，这个goroutine向里面写数据，函数要返回这个通道。
// 有1个函数来组织流水线，我们例子中是Consumer函数

// func producer(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for _, n := range nums {
// 			out <- n
// 		}
// 	}()
// 	return out
// }

func producer(nums int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < nums; i++ {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()
	return out
}

func Consumer() {
	in := producer(10000000)
	ch := square(in)

	//consumer
	for _ = range ch {
		// fmt.Printf("%d ", ret)
	}
	fmt.Println()
}
