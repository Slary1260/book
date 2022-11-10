/*
 * @Author: tj
 * @Date: 2022-11-10 14:13:25
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 16:06:04
 * @FilePath: \book\main.go
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
	"book/database/impl"
	"syscall/js"
)

func main() {
	m, err := impl.GetMemoryDb()
	if err != nil {
		panic(err)
	}

	done := make(chan int, 0)
	js.Global().Set("setData", js.FuncOf(m.SetData))
	js.Global().Set("readData", js.FuncOf(m.ReadData))
	<-done
}
