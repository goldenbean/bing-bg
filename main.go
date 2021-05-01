package main

import (
	"flag"
	"fmt"
	"os"
)

var savePath = ""

func init() {
	home := os.Getenv("HOME")
	flag.StringVar(&savePath, "o", home, "save path")
	flag.Parse()

	fmt.Println("user home: ", home)
	fmt.Printf("save path: %s\n", savePath)
}

func main() {
	bingUrl := "https://www.bing.com/?mkt=zh-CN"

	ret := fetchImageUrl(bingUrl)

	for _, txt := range ret {
		url := "https://cn.bing.com/th?id=" + txt
		httpGetBinaryToFile(url, savePath+"/"+txt)
	}

}
