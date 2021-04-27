package main

import (
	"fmt"
	"log"

	registry "golang.org/x/sys/windows/registry"
)

func main() {
	//k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Personalization`, registry.QUERY_VALUE)
	key, exists, err := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\PersonalizationCSP`, registry.ALL_ACCESS)

	if err != nil {
		log.Fatal(err)
	}
	//defer k.Close()
	defer key.Close()

	if exists {
		println(`键已存在`)
	} else {
		println(`新建注册表键`)
	}

	// 写入：字符串
	key.SetStringValue(`LockScreenImagePath`, `C:\Users\liuyan\Desktop\11.jpg`)
	key.SetStringValue(`LockScreenImageUrl`, `C:\Users\liuyan\Desktop\11.jpg`)
	// 写入：32位整形值
	key.SetDWordValue(`LockScreenImageStatus`, uint32(1))

	s, _, err := key.GetStringsValue("LockScreenImagePath")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
