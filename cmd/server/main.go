/*
 * @Author: tj
 * @Date: 2022-11-10 16:13:13
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 17:28:27
 * @FilePath: \book\cmd\server\main.go
 */
package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":9090", http.FileServer(http.Dir("../../static")))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
