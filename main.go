package main

import (
	"log"
)

func main() {
	if App.UseFileLog {
		logFile := initLog(App.LogFile)
		defer logFile.Close()
	}
	urlPipe := make(chan string)
	if App.UseHttp {
		initHttp(urlPipe)
	}
	if App.UseFiles {
		initFiles(urlPipe)
	}
	log.Println("urlConstantOpener is running...")
	for {
		select {
		case url := <-urlPipe:
			openUrl(url)
		}
	}
}
