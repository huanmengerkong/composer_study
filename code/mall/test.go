package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/aa",GetInfo)
	fmt.Println("这个东西起来了")
	http.ListenAndServe("0.0.0.0:8000",nil)
}
func  GetInfo(r http.ResponseWriter ,w *http.Request)  {
	fmt.Println(2222)
}
