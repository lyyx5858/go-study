package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	op1,err1 :=os.Open("/Users/liuyan/go/src/go-study/IO.COPY/a1.exe")
	op2,err2 :=os.Open("/Users/liuyan/go/src/go-study/IO.COPY/a2.exe")
	of1,err3 :=os.Create("/Users/liuyan/go/src/go-study/IO.COPY/a3.exe")
	of2,err4 :=os.Create("/Users/liuyan/go/src/go-study/IO.COPY/a4.exe")

	if err1!=nil || err2!=nil||err3!=nil||err4!=nil{
		fmt.Println("文件拷贝失败",err1,err2,err3)
		return
	}
	defer op1.Close()
	defer op2.Close()
	defer of1.Close()
	defer of2.Close()

	start:=time.Now()
	b:=make([]byte,64*1024)
	io.CopyBuffer(of1,op1,b)
	end:=time.Now()
	fmt.Println(end.Sub(start))

	s1:=time.Now()
	c:=make([]byte,1*1024)
	io.CopyBuffer(of2,op2,c)
	e1:=time.Now()
	fmt.Println(e1.Sub(s1))

}
