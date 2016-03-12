package main

import (
	"time"
)

func initFiles(pipe chan string) {
	go func(pipe chan string) {
		for {
			time.Sleep(1 * time.Second)
			checkDirectory(pipe)
		}
	}(pipe)
}

func checkDirectory(pipe chan string) {
	files := readDirectory(App.DirectoryToScan)
	length := len(files)
	for i := 0; i < length; i++ {
		url := getUrlFromFile(files[i], App.DirectoryToScan)
		if url != "" {
			pipe <- url
		}
	}
}
