/*
 * @Author: tj
 * @Date: 2022-11-10 17:51:30
 * @LastEditors: tj
 * @LastEditTime: 2022-11-11 15:43:18
 * @FilePath: \book\database\impl\index.go
 */
package impl

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/buntdb"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (m *MemoryDb) GetIndexLess(this js.Value, args []js.Value) interface{} {
	m.lessMutex.RLock()
	defer m.lessMutex.RUnlock()

	data, err := json.Marshal(&m.indexLessMap)
	if err != nil {
		fmt.Println("GetIndexLess Marshal error:", err.Error())
		return js.ValueOf(err.Error())
	}

	return js.ValueOf(string(data))
}

// 创建索引 至少三个参数:索引key，索引模式，索引排序类型(可自定义)，第四个参数为json的排序类型(可以多个,以","分隔)
// '*' matches on any number characters
// '?' matches on any one character
func (m *MemoryDb) AddIndex(this js.Value, args []js.Value) interface{} {
	// if len(args) < 3 {
	// 	return os.ErrInvalid.Error()
	// }

	lessSeq, err := strconv.Atoi(args[2].String())
	if err != nil {
		fmt.Println("AddIndex error:", err.Error())
		return js.ValueOf(err.Error())
	}

	switch lessSeq {
	case 0: // IndexJSON
		jsonLessStr := args[3].String()
		lessFuns := strings.Split(jsonLessStr, ",")

		if len(lessFuns) == 1 {
			// 单个
			err := m.db.CreateIndex(args[0].String(), args[1].String(), buntdb.IndexJSON(lessFuns[0]))
			if err != nil {
				return js.ValueOf(err.Error())
			}
		} else {
			// 多个
			funs := make([]func(a, b string) bool, 0, 4)
			for _, v := range lessFuns {
				funs = append(funs, buntdb.IndexJSON(v))
			}

			err := m.db.CreateIndex(args[0].String(), args[1].String(), funs...)
			if err != nil {
				return js.ValueOf(err.Error())
			}
		}

	case 1: // IndexBinary
		err = m.db.CreateIndex(args[0].String(), args[1].String(), buntdb.IndexBinary)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 2: // IndexFloat
		err = m.db.CreateIndex(args[0].String(), args[1].String(), buntdb.IndexFloat)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 3: // IndexInt
		err = m.db.CreateIndex(args[0].String(), args[1].String(), buntdb.IndexInt)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 4: // IndexString
		err = m.db.CreateIndex(args[0].String(), args[1].String(), buntdb.IndexString)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 5: // IndexUint
		err = m.db.CreateIndex(args[0].String(), args[1].String(), buntdb.IndexUint)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	default:
		fmt.Println("unknow lessSeq:" + strconv.Itoa(lessSeq))
		return js.ValueOf("unknow lessSeq:" + strconv.Itoa(lessSeq))
	}

	m.mutex.Lock()
	m.indexMap[args[0].String()] = buntdb.IndexString
	m.patternMap[args[0].String()] = args[1].String()
	m.mutex.Unlock()

	return js.ValueOf("")
}

// 替换索引 至少三个参数:索引key，索引模式，索引排序类型(可自定义)，第四个参数为json的排序类型(可以多个,以","分隔)
// '*' matches on any number characters
// '?' matches on any one character
func (m *MemoryDb) ReplaceIndex(this js.Value, args []js.Value) interface{} {
	// if len(args) < 3 {
	// 	return os.ErrInvalid.Error()
	// }

	lessSeq, err := strconv.Atoi(args[2].String())
	if err != nil {
		fmt.Println("AddIndex error:", err.Error())
		return js.ValueOf(err.Error())
	}

	switch lessSeq {
	case 0: // IndexJSON
		jsonLessStr := args[3].String()
		lessFuns := strings.Split(jsonLessStr, ",")

		if len(lessFuns) == 1 {
			// 单个
			err := m.db.ReplaceIndex(args[0].String(), args[1].String(), buntdb.IndexJSON(lessFuns[0]))
			if err != nil {
				return js.ValueOf(err.Error())
			}
		} else {
			// 多个
			funs := make([]func(a, b string) bool, 0, 4)
			for _, v := range lessFuns {
				funs = append(funs, buntdb.IndexJSON(v))
			}

			err := m.db.ReplaceIndex(args[0].String(), args[1].String(), funs...)
			if err != nil {
				return js.ValueOf(err.Error())
			}
		}

	case 1: // IndexBinary
		err = m.db.ReplaceIndex(args[0].String(), args[1].String(), buntdb.IndexBinary)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 2: // IndexFloat
		err = m.db.ReplaceIndex(args[0].String(), args[1].String(), buntdb.IndexFloat)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 3: // IndexInt
		err = m.db.ReplaceIndex(args[0].String(), args[1].String(), buntdb.IndexInt)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 4: // IndexString
		err = m.db.ReplaceIndex(args[0].String(), args[1].String(), buntdb.IndexString)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	case 5: // IndexUint
		err = m.db.ReplaceIndex(args[0].String(), args[1].String(), buntdb.IndexUint)
		if err != nil {
			return js.ValueOf(err.Error())
		}

	default:
		fmt.Println("unknow lessSeq:" + strconv.Itoa(lessSeq))
		return js.ValueOf("unknow lessSeq:" + strconv.Itoa(lessSeq))
	}

	m.mutex.Lock()
	m.indexMap[args[0].String()] = buntdb.IndexString
	delete(m.patternMap, args[0].String())
	m.patternMap[args[0].String()] = args[1].String()
	m.mutex.Unlock()

	return js.ValueOf("")
}

func (m *MemoryDb) DeleteIndex(this js.Value, args []js.Value) interface{} {
	// if len(args) != 1 {
	// 	return os.ErrInvalid.Error()
	// }

	err := m.db.DropIndex(args[0].String())
	if err != nil {
		fmt.Println("DeleteIndex error:", err.Error())
		return err.Error()
	}

	m.mutex.Lock()
	delete(m.indexMap, args[0].String())
	delete(m.patternMap, args[0].String())
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
