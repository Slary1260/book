/*
 * @Author: tj
 * @Date: 2022-11-10 14:13:25
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 16:45:49
 * @FilePath: \book\cmd\wasm\main.go
 */
/*
 * @Author: tj
 * @Date: 2022-11-10 14:13:25
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 14:13:31
 * @FilePath: \book\main.go
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
