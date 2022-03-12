package main

import (
	"flag"
	"fmt"
	"os"
)

var savePath = ""

func init() {
	home := os.Getenv("HOME")
	flag.StringVar(&savePath, "o", "./", "save path")
	flag.Parse()

	fmt.Println("user home: ", home)
	fmt.Printf("save path: %s\n", savePath)
}

func main() {
	bingUrl := "https://www.bing.com/?mkt=zh-CN"

	ret := fetchImageUrl(bingUrl)

	fmt.Println(ret)

	for _, txt := range ret {
		url := "https://s.cn.bing.net/th?id=" + txt
		httpGetBinaryToFile(url, savePath+"/"+txt)
	}

}
