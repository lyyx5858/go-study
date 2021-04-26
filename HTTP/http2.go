package main

import (
	"fmt"

	"io/ioutil"

	"log"

	"net/http"
)

func main() {

	resp, err := http.Get("http://www.baidu.com")

	if err != nil {

		log.Fatal(err)

	}

	bData, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Fatal(err)

	}

	defer resp.Body.Close()

	fmt.Println(string(bData))

}
