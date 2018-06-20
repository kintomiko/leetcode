package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			http.Get("http://localhost:4000")
			fmt.Println("finished call")
		}()
	}
	fmt.Println("created 1000 calls")
	for {
		fmt.Println("total goroutine number: " + strconv.Itoa(runtime.NumGoroutine()))
		time.Sleep(time.Second)
	}
}
