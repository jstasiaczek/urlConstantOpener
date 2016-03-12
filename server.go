package main

import (
	"net/http"
	"strconv"
)

func initHttp(passer chan string) {
	go func(passer chan string) {
		responseHandler := &DataPasser{passer: passer}
		http.HandleFunc("/openUrl", responseHandler.HandleUrl)
		url := ":" + strconv.FormatInt(int64(App.HttpPort), 10)
		http.ListenAndServe(url, nil)
	}(passer)
}

type DataPasser struct {
	passer chan string
}

func (p *DataPasser) HandleUrl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form == nil || r.Header.Get("secret") != App.SecretKey {
		w.Write([]byte("Fail!"))
		return
	}

	p.passer <- r.Form.Get("url")
	w.Write([]byte("OK"))
}
