package main

import (
	"context"
	"fmt"

	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)   //注意返回的cancel是个FUNC
	go test(ctx)  //因为CTX是有Cancel功能的context，就如同给协程test加装了一个开关，用cancel就可以取消协和f
	time.Sleep(time.Second * 3)
	cancel() //等3秒就取消test协程，但cancel只是传递一个信号给test，但这个信号协程接收到如何处理，由协程自己定义。
			//这个信号就是将ctx.Done这个channel赋一个值。
	time.Sleep(time.Second * 1)
}

func test(ctx context.Context) {
	i := 1

	var x <- chan int  // x是接收整形的channel
	x=make(chan int) //x 是双向的
	x=make(<- chan int) //x是单向的。

	var y chan <- int  // y是可发送整形的channel, 利用var语句定义是没有初始过的。

	var x1 chan struct{} //双向的
	var z <- chan struct{}  //z是可接收结构的channel
	z=make(<-chan struct{})

	fmt.Println(x,y,x1,z)

	fmt.Println(x)
	fmt.Println(<-x)   //注意这个和上面的区别，有了这个箭头，说明读操作。读不到就阻塞。



	for {
		z = ctx.Done()
		fmt.Println(z)

		select {
		case <-ctx.Done():
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


