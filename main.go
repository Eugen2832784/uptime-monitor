package main

import (
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}
func main(){
	check("https://go.dev")
	check("https://github.com")
	check("https://this-site-does-not-exist.example")
	check("https://httpbin.org/status/500")
}

func check(url string){
	start := time.Now()
	resp,err:= client.Get(url)
	took := fmt.Sprintf("%dms", time.Since(start).Milliseconds())
	if err!= nil{
		fmt.Printf("%-5s %3s  %7s  %s  (%v)\n", "DOWN", "---", "---", url, err)
		return
	}

	defer resp.Body.Close()

	status := "DOWN"
	if resp.StatusCode>= 200 && resp.StatusCode<400{
		status = "UP"
	}
	fmt.Printf("%-5s %3d  %7s  %s\n", status, resp.StatusCode, took, url)
}