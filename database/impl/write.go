/*
 * @Author: tj
 * @Date: 2022-11-10 14:59:19
 * @LastEditors: tj
 * @LastEditTime: 2022-11-14 10:34:59
 * @FilePath: \book\database\impl\write.go
 */
package impl

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"

	"github.com/tidwall/buntdb"
)

// 设置数据
// key,value,过期毫秒数
func (m *MemoryDb) SetData(this js.Value, args []js.Value) interface{} {
	// if len(args) != 3 {
	// 	return os.ErrInvalid.Error()
	// }

	// TODO this is demo
	if len(args) == 0 {
		err := m.db.CreateIndex("last_name", "*", buntdb.IndexJSON("name.last"))
		if err != nil {
			return js.ValueOf(err.Error())
		}

		err = m.db.CreateIndex("age", "*", buntdb.IndexJSON("age"))
		if err != nil {
			return js.ValueOf(err.Error())
		}

		err = m.db.CreateIndex("address", "*", buntdb.IndexString)
		if err != nil {
			return js.ValueOf(err.Error())
		}

		err = m.db.Update(func(tx *buntdb.Tx) error {
			tx.Set("1", `{"name":{"first":"Tom","last":"Johnson"},"age":38}`, nil)
			tx.Set("2", `{"name":{"first":"Janet","last":"Prichard"},"age":47}`, nil)
			tx.Set("3", `{"name":{"first":"Carol","last":"Anderson"},"age":52}`, nil)
			tx.Set("4", `{"name":{"first":"Alan","last":"Cooper"},"age":28}`, nil)
			tx.Set("5", `"address":28`, nil)
			tx.Set("6", `"address":15`, nil)
			return nil
		})
		if err != nil {
			fmt.Println("SetData Update error:", err.Error())
			return js.ValueOf(err.Error())
		}

		return js.ValueOf("")
	}

	opts := &buntdb.SetOptions{}
	if args[2].String() != "" {
		count, err := strconv.Atoi(args[2].String())
		if err != nil {
			fmt.Println("SetData Atoi error:", err.Error())
			return js.ValueOf(err.Error())
		}
		opts.Expires = true
		opts.TTL = time.Millisecond * time.Duration(count)
	}

	err := m.db.Update(func(tx *buntdb.Tx) error {
		previousValue, replaced, err := tx.Set(args[0].String(), args[1].String(), opts)
		if err != nil {
			fmt.Println("SetData error:", err.Error())
			return err
		}
		fmt.Println("replaced:", replaced)
		fmt.Println("previousValue:", previousValue)

		return nil
	})
	if err != nil {
		fmt.Println("SetData Update error:", err.Error())
		return js.ValueOf(err.Error())
	}

	return js.ValueOf("")
}

// 删除数据
// 通过索引模式获取key，然后再根据key删除数据
// 索引模式(空:*),删除的key条件,删除的value条件
func (m *MemoryDb) DeleteKeyByPattern(this js.Value, args []js.Value) interface{} {
	// if len(args) != 3 {
	// 	return os.ErrInvalid.Error()
	// }

	pattern := args[0].String()
	find := false
	m.mutex.RLock()
	for _, v := range m.patternMap {
		if v == pattern {
			find = true
			break
		}
	}
	m.mutex.RUnlock()

	if !find {
		return js.ValueOf("")
	}

	if pattern == "" {
		pattern = "*"
	}

	delkeys := make([]string, 0, 4)
	err := m.db.Update(func(tx *buntdb.Tx) error {
		tx.AscendKeys(pattern, func(k, v string) bool {
			// TODO 过滤条件
			if m.someCondition(args[1].String(), args[2].String(), k, v) == true {
				delkeys = append(delkeys, k)
			}
			return true
		})

		for _, k := range delkeys {
			_, err := tx.Delete(k)
			if err != nil {
				fmt.Println("DeleteKeyByPattern Delete error:", err.Error())
				return err
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println("DeleteKeyByPattern Update error:", err.Error())
		return js.ValueOf(err.Error())
	}

	return js.ValueOf("")
}

func (m *MemoryDb) someCondition(keyCondition, valueCondition, key, value string) bool {
	// TODO 删除条件
	return false
}

func (m *MemoryDb) DeleteAll(this js.Value, args []js.Value) interface{} {
	err := m.db.Update(func(tx *buntdb.Tx) error {
		err := tx.DeleteAll()
		if err != nil {
			fmt.Println("DeleteAll error:", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		fmt.Println("DeleteByKey Update error:", err.Error())
		return js.ValueOf(err.Error())
	}

	return js.ValueOf("")
}
