/*
 * @Author: tj
 * @Date: 2022-11-10 14:59:28
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 17:58:10
 * @FilePath: \book\database\impl\read.go
 */
package impl

import (
	"fmt"
	"log"
	"strings"
	"syscall/js"

	// jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/buntdb"
)

func (m *MemoryDb) ReadData(this js.Value, args []js.Value) interface{} {
	// if len(args) != 1 {
	// 	return os.ErrInvalid.Error()
	// }

	ret := make([]string, 0)
	m.db.View(func(tx *buntdb.Tx) error {
		fmt.Println("Order by last name")
		datail := make([]string, 0)
		tx.Ascend("last_name", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)
		ret = append(ret, "******************")

		fmt.Println("Order by age")
		datail = make([]string, 0)
		tx.Ascend("age", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)
		ret = append(ret, "******************")

		fmt.Println("Order by age range 30-50")
		datail = make([]string, 0)
		tx.AscendRange("age", `{"age":30}`, `{"age":50}`, func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)
		ret = append(ret, "******************")

		fmt.Println("get")
		datail = make([]string, 0)
		value, err := tx.Get("4")
		if err != nil {
			ret = append(ret, err.Error())
		} else {
			ret = append(ret, value)
		}
		fmt.Printf("%d: %s\n", 4, value)
		ret = append(ret, "******************")

		fmt.Println("Order by name")
		datail = make([]string, 0)
		tx.Ascend("name", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)
		ret = append(ret, "******************")

		// 没有加index 也查询出来了
		fmt.Println("Order by age")
		datail = make([]string, 0)
		tx.Ascend("age", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)
		ret = append(ret, "******************")

		// 没有的key查不到数据
		fmt.Println("Order by count")
		datail = make([]string, 0)
		tx.Ascend("count", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)
		ret = append(ret, "******************")

		fmt.Println("indexs")
		indexs, _ := tx.Indexes()
		ret = append(ret, indexs...)
		ret = append(ret, "******************")

		return nil
	})

	return js.ValueOf(strings.Join(ret, ","))
}

// 通过KEY获取1个数据
func (m *MemoryDb) GetDataByKey(this js.Value, args []js.Value) interface{} {
	// if len(args) != 1 {
	// 	return os.ErrInvalid.Error()
	// }

	result := ""
	err := m.db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get(args[0].String())
		if err != nil {
			fmt.Println("GetDataByKey error:", err.Error())
			return err
		}

		result = value

		return nil
	})
	if err != nil {
		fmt.Println("GetDataByKey View error:", err.Error())
		return err.Error()
	}

	return js.ValueOf(result)
}

// 通过索引获取多个数据
func (m *MemoryDb) GetAllDataByIndex(this js.Value, args []js.Value) interface{} {
	// if len(args) != 1 {
	// 	return os.ErrInvalid.Error()
	// }

	ret := make([]string, 0)
	err := m.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(args[0].String(), func(key, value string) bool {
			log.Printf("%s: %s\n", key, value)
			ret = append(ret, value)
			return true
		})
		if err != nil {
			fmt.Println("GetAllDataByIndex error:", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("GetAllDataByIndex View error:", err.Error())
		return err.Error()
	}

	return js.ValueOf(strings.Join(ret, ","))
}

// 通过索引的范围获取多个数据 需要三个参数
func (m *MemoryDb) GetAllDataByRange(this js.Value, args []js.Value) interface{} {
	// if len(args) != 3 {
	// 	return os.ErrInvalid.Error()
	// }

	ret := make([]string, 0)
	err := m.db.View(func(tx *buntdb.Tx) error {
		key := args[0].String()
		min := `{` + args[0].String() + `:` + args[1].String() + `}`
		max := `{` + args[0].String() + `:` + args[2].String() + `}`
		err := tx.AscendRange(key, min, max, func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			ret = append(ret, value)
			return true
		})
		if err != nil {
			fmt.Println("GetAllDataByRange error:", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("GetAllDataByRange View error:", err.Error())
		return err.Error()
	}

	return js.ValueOf(strings.Join(ret, ","))
}

// 通过索引遍历所有数据
func (m *MemoryDb) GetAllData(this js.Value, args []js.Value) interface{} {
	// if len(args) != 1 {
	// 	return os.ErrInvalid.Error()
	// }

	ret := make([]string, 0)
	err := m.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend(args[0].String(), func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			ret = append(ret, value)
			return true
		})
		if err != nil {
			fmt.Println("GetAllData error:", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("GetAllData View error:", err.Error())
		return err.Error()
	}

	return js.ValueOf(strings.Join(ret, ","))
}

func (m *MemoryDb) GetDataLength(this js.Value, args []js.Value) interface{} {
	result := 0
	err := m.db.View(func(tx *buntdb.Tx) error {
		count, err := tx.Len()
		if err != nil {
			fmt.Println("GetDataLength error:", err.Error())
			return err
		}

		result = count

		return nil
	})
	if err != nil {
		fmt.Println("GetDataLength View error:", err.Error())
		return err.Error()
	}

	return js.ValueOf(result)
}
