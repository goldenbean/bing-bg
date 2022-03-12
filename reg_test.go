package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// go test -v -run TestParseImageUrl reg_test.go reg.go
func TestParseImageUrl(t *testing.T) {
	t.Log("TestParseImageUrl")

	content := readFile("reg_test.txt")
	fmt.Println(content)

	imageUrls := ParseImageUrl(content)
	fmt.Println("len: ", len(imageUrls))
	fmt.Println(imageUrls)
	for _, url := range imageUrls {
		fmt.Println(extractImageFilename(url))
	}

}
