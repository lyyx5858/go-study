package main

import (
	"context"
"fmt"
"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	go f(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 1)
}

func f(ctx context.Context) {
	i := 1
	for {
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


