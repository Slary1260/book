/*
 * @Author: tj
 * @Date: 2022-11-10 14:13:25
 * @LastEditors: tj
 * @LastEditTime: 2022-11-11 15:43:56
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

	// TODO this is demo
	js.Global().Set("readData", js.FuncOf(m.ReadData))

	// 通过KEY获取1个数据
	js.Global().Set("getDataByKey", js.FuncOf(m.GetDataByKey))
	// 通过索引获取多个数据
	js.Global().Set("getAllDataByIndex", js.FuncOf(m.GetAllDataByIndex))
	// 通过索引的范围获取多个数据 需要三个参数
	js.Global().Set("getAllDataByRange", js.FuncOf(m.GetAllDataByRange))
	// 通过索引遍历所有数据
	js.Global().Set("getAllData", js.FuncOf(m.GetAllData))
	js.Global().Set("getDataLength", js.FuncOf(m.GetDataLength))

	// key,value,过期毫秒数
	js.Global().Set("setData", js.FuncOf(m.SetData))
	// 索引模式(空:*),删除的key条件,删除的value条件
	js.Global().Set("deleteKeyByPattern", js.FuncOf(m.DeleteKeyByPattern))
	js.Global().Set("deleteAll", js.FuncOf(m.DeleteAll))

	js.Global().Set("getIndexLess", js.FuncOf(m.GetIndexLess))
	// 创建索引 至少三个参数:索引key，索引模式，索引排序类型(可自定义)，第四个参数为json的排序类型(可以多个,以","分隔)
	// '*' matches on any number characters
	// '?' matches on any one character
	js.Global().Set("addIndex", js.FuncOf(m.AddIndex))
	// 替换索引 至少三个参数:索引key，索引模式，索引排序类型(可自定义)，第四个参数为json的排序类型(可以多个,以","分隔)
	js.Global().Set("replaceIndex", js.FuncOf(m.ReplaceIndex))
	js.Global().Set("deleteIndex", js.FuncOf(m.DeleteIndex))
	// 获取所有的索引
	js.Global().Set("getAllIndex", js.FuncOf(m.GetAllIndex))

	<-make(chan bool)
}
