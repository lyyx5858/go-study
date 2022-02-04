package main

import (
	"context"
	"fmt"

	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx) //注意返回的cancel是个FUNC
	go test(ctx)                           //因为CTX是有Cancel功能的context，就如同给协程test加装了一个开关，用cancel就可以取消协和f
	time.Sleep(time.Second * 3)
	cancel() //等3秒就取消test协程，但cancel只是传递一个信号给test，但这个信号协程接收到如何处理，由协程自己定义。
	// 这个函数的作用其实就是将ctx.Done这个channel赋一个值。
	time.Sleep(time.Second * 1)
}

func test(ctx context.Context) {
	i := 1
	for {
		z := ctx.Done()
		fmt.Println(z)

		select {
		case <-ctx.Done(): //当主程序里使用cancel()这句话后，这个值就可以读到了，说明主程序指示：该结束了。
			fmt.Println("DONE")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
