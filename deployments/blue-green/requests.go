package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		request()
		time.Sleep(2 * time.Second)
	}
}

func request() {
	resp, err := http.Get("http://localhost:30000")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Print(string(body))
}
