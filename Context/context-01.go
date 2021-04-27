package main

import (
	"context"
	"fmt"
	"time"
)

func main(){

	ctx:=context.WithValue(context.Background(),"begin","从台词看到一句话")

	go g(ctx)
	time.Sleep(time.Second)

}


func g(ctx context.Context){
	fmt.Println(ctx.Value("begin"))
	fmt.Println("你是猪")
}