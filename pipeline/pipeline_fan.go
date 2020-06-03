package pipeline

import (
	"fmt"
	"sync"
)

func merge(cs ...<-chan int) <-chan int {
	//参考链接：https://mp.weixin.qq.com/s?__biz=Mzg3MTA0NDQ1OQ==&mid=2247483680&idx=1&sn=de463ebbd088c0acf6c2f0b5f179f38d&scene=21#wechat_redirect
	/* 测试报告：
	同为10000000条消息
	1、pipeline_use.go
	go run main.go 2  9.03s user 3.82s system 184% cpu 6.963 total

	2、pipeline_fan.go
	1）merge中的out channel为无缓冲通道
	go run main.go 1  19.37s user 3.85s system 317% cpu 7.312 total

	2）merge中的out channel为100缓冲通道
	go run main.go 1  16.82s user 3.04s system 307% cpu 6.469 total

	*/
	out := make(chan int, 100)
	var wg sync.WaitGroup
	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go collect(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func ConsumerFan() {
	in := producer(10000000)

	//FAN-OUT模式：多个goroutine从同一个通道读取数据，直到该通道关闭。（扇出，用来分发任务）
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	//FAN-IN模式：1个goroutine从多个通道读取数据，直到这些通道关闭。（扇入，用来收集处理的结果）
	for _ = range merge(c1, c2, c3) {
		// fmt.Printf("%3d", ret)
	}
	fmt.Println()

}
