package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/toqueteos/webbrowser"
)

// read directory and return content or empty array on fail
func readDirectory(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
		return []os.FileInfo{}
	}
	return files
}

// init log file
func initLog(path string) *os.File {
	_, err := os.Stat(path)
	if err != nil {
		createLogFile(path)
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	return f
}

// create log file
func createLogFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	f.Close()
}

// read file and return it content or empty string on fail
func readFileFailsafe(path string) string {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return strings.Trim(string(raw), "\n\r ")
}

func getUrlFromFile(file os.FileInfo, directoryPath string) string {
	if file.IsDir() {
		return ""
	}
	path := directoryPath + "/" + file.Name()
	url := readFileFailsafe(path)
	// sometimes somehow we are reading file when it is created but without content
	if url != "" {
		os.Remove(path)
	}
	return url
}

// try to open url from file
func openUrl(url string) {
	if url == "" {
		return
	}
	err := webbrowser.Open(url)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Opened url: " + url)
}
