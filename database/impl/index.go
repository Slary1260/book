/*
 * @Author: tj
 * @Date: 2022-11-10 17:51:30
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 18:02:24
 * @FilePath: \book\database\impl\index.go
 */
package impl

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/tidwall/buntdb"
)

// 创建索引 三个参数:索引key，索引模式，索引排序类型(可自定义)
// '*' matches on any number characters
// '?' matches on any one character
func (m *MemoryDb) AddIndex(this js.Value, args []js.Value) interface{} {
	// if len(args) != 3 {
	// 	return os.ErrInvalid.Error()
	// }

	// TODO 索引排序类型
	err := m.db.CreateIndex(args[0].String(), args[1].String(), buntdb.IndexString)
	if err != nil {
		return js.ValueOf(err.Error())
	}

	m.mutex.Lock()
	m.indexMap[args[0].String()] = buntdb.IndexString
	m.mutex.Unlock()

	return js.ValueOf("")
}

func (m *MemoryDb) ReplaceIndex(this js.Value, args []js.Value) interface{} {
	// if len(args) != 3 {
	// 	return os.ErrInvalid.Error()
	// }

	// TODO 索引排序类型
	err := m.db.ReplaceIndex(args[0].String(), args[1].String(), buntdb.IndexString)
	if err != nil {
		fmt.Println("ReplaceIndex error:", err.Error())
		return err.Error()
	}

	m.mutex.Lock()
	m.indexMap[args[0].String()] = buntdb.IndexString
	m.mutex.Unlock()

	return js.ValueOf("")
}

// 获取所有的索引
func (m *MemoryDb) GetAllIndex(this js.Value, args []js.Value) interface{} {
	ret, err := m.db.Indexes()
	if err != nil {
		fmt.Println("GetAllIndex error:", err.Error())
		return err.Error()
	}

	return js.ValueOf(strings.Join(ret, ","))
}
