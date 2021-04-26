package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func readRaw(path string) string {
	start := time.Now()
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	defer func() {
		fi.Close()
		fmt.Printf("[readRaw] cost time %v \n", time.Now().Sub(start))
	}()

	var data []byte
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		data = append(data, buf[:n]...)
		if 0 == n {
			break
		}
	}
	return string(data)
}

func readWithBufferIO(path string) string {
	start := time.Now()
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		fi.Close()
		fmt.Printf("[readWithBufferIO] cost time %v \n", time.Now().Sub(start))
	}()

	r := bufio.NewReader(fi)

	var data []byte

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		data = append(data, buf[:n]...)
	}
	return string(data)
}

func readWithIOUtil(path string) string {
	start := time.Now()
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		fi.Close()
		fmt.Printf("[readWithIOUtil] cost time %v \n", time.Now().Sub(start))
	}()

	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func main() {
	file := "test.txt"

	readRaw(file)
	readWithBufferIO(file)
	readWithIOUtil(file)
}
