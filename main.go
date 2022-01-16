/**
 *@Author luojunying
 *@Date 2022-01-16 19:35
 */
package main

import (
	"Xy/framework"
	"fmt"
	"net/http"
)

func main() {
	core := framework.NewCore()
	server := http.Server{
		Addr:    ":8888",
		Handler:  core,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf(err.Error())
	}
}
