package main

import (
	"fmt"
	"github.com/e9571/marmot/miner"
)

func postFile(filename string, targetUrl string) {
	worker, _ := miner.New(nil)
	result, err := worker.SetUrl(targetUrl).SetBData([]byte("dddd")).SetFileInfo(filename+".xxxx", "uploadfile").PostFILE()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(result))
	}
}

// sample usage
func main() {
	targetUrl := "http://127.0.0.1:1789/upload"
	filename := "./doc.go"
	postFile(filename, targetUrl)
}
