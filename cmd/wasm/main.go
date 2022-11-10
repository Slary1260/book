/*
 * @Author: tj
 * @Date: 2022-11-10 14:13:25
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 16:53:45
 * @FilePath: \book\cmd\wasm\main.go
 */
package main

import (
	"syscall/js"

	"book/database/impl"
)

func main() {
	m, err := impl.GetMemoryDb()
	if err != nil {
		panic(err)
	}

	js.Global().Set("setData", js.FuncOf(m.SetData))
	js.Global().Set("readData", js.FuncOf(m.ReadData))

	<-make(chan bool)
}
