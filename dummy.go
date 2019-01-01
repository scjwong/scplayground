package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var count int

func helloWorld(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintf(w, "Hello World! %d\n", count)
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Body len: %d\n", len(body))
	fmt.Printf("Wrote 'Hello World! %d\n", count)
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
