package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func fetchImageUrl(url string) []string {
	str := httpGet(url)

	// 非贪婪模式
	reg := regexp.MustCompile(`background-image:url\(\/th.*?\)`)
	imageUrls := reg.FindAllString(str, -1)

	ret := []string{}
	for _, txt := range imageUrls {
		fmt.Println(txt)
		reg2 := regexp.MustCompile(`=.*?&`)

		for _, param := range reg2.FindStringSubmatch(txt) {
			ret = append(ret, param[1:len(param)-1])
			continue
		}
	}
	return ret
}

func httpGet(url string) string {

	res, err := http.Get(url)
	if err != nil {
		return ""
		//panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
		// handle error
	}

	// fmt.Println(string(body))
	//	fmt.Println(res.Status)

	return string(body)
}

func httpGetBinaryToFile(url string, path string) (string, bool) {

	fmt.Printf("url: %s \npath: [%s]\n", url, path)

	res, err := http.Get(url)
	if err != nil {
		return err.Error(), false
	}
	defer res.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return "", true
}
