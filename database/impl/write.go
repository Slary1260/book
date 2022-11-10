/*
 * @Author: tj
 * @Date: 2022-11-10 14:59:19
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 18:39:53
 * @FilePath: \book\database\impl\write.go
 */
package impl

import (
	"fmt"
	"syscall/js"

	"github.com/tidwall/buntdb"
)

func (m *MemoryDb) SetData(this js.Value, args []js.Value) interface{} {
	// if len(args) != 2 {
	// 	return os.ErrInvalid.Error()
	// }

	err := m.db.CreateIndex("last_name", "*", buntdb.IndexJSON("name.last"))
	if err != nil {
		return js.ValueOf(err.Error())
	}

	err = m.db.CreateIndex("age", "*", buntdb.IndexJSON("age"))
	if err != nil {
		return js.ValueOf(err.Error())
	}

	// err = m.db.CreateIndex("name", "*", buntdb.IndexString)
	// if err != nil {
	// 	return js.ValueOf(err.Error())
	// }

	err = m.db.Update(func(tx *buntdb.Tx) error {
		tx.Set("1", `{"name":{"first":"Tom","last":"Johnson"},"age":38}`, nil)
		tx.Set("2", `{"name":{"first":"Janet","last":"Prichard"},"age":47}`, nil)
		tx.Set("3", `{"name":{"first":"Carol","last":"Anderson"},"age":52}`, nil)
		tx.Set("4", `{"name":{"first":"Alan","last":"Cooper"},"age":28}`, nil)
		return nil
	})
	if err != nil {
		return js.ValueOf(err.Error())
	}

	return js.ValueOf("")
}

func (m *MemoryDb) DeleteByKey(this js.Value, args []js.Value) interface{} {
	// if len(args) != 1 {
	// 	return os.ErrInvalid.Error()
	// }

	// TODO
	err := m.db.Update(func(tx *buntdb.Tx) error {
		// tx.AscendKeys("object:*", func(k, v string) bool {
		// 	if someCondition(k) == true {
		// 		delkeys = append(delkeys, k)
		// 	}
		// 	return true // continue
		// })

		return nil
	})
	if err != nil {
		fmt.Println("DeleteByKey error:", err.Error())
		return js.ValueOf(err.Error())
	}

	return js.ValueOf("")
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
		fmt.Println("DeleteByKey error:", err.Error())
		return js.ValueOf(err.Error())
	}

	return js.ValueOf("")
}
