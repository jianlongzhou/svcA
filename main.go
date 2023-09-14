package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/service/a", helloHandler)
	fmt.Println("start service: A")

	go func() {
		for {
			request()
			time.Sleep(time.Second)
		}
	}()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("start service err: ", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get request header: ", r.Header)
	fmt.Fprint(w, "A, Hello, World! version: ", os.Getenv("VERSION"))
}

func request() {
	host := os.Getenv("HOST")
	resp, err := http.Get(fmt.Sprintf("%s:8081/service/b", host))
	if err != nil {
		fmt.Println("call service b err: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read service rsp body err: ", err)
		return
	}
	fmt.Println("svcA -> svcB: ", string(body))
}
