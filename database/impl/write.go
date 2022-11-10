/*
 * @Author: tj
 * @Date: 2022-11-10 14:59:19
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 15:29:47
 * @FilePath: \book\database\impl\write.go
 */
package impl

import (
	"syscall/js"

	"github.com/tidwall/buntdb"
)

func (m *MemoryDb) SetData(this js.Value, args []js.Value) interface{} {
	err := m.db.CreateIndex("names", "*", buntdb.IndexString)
	if err != nil {
		return js.ValueOf(err.Error())
	}

	err = m.db.Update(func(tx *buntdb.Tx) error {
		tx.Set("user:0:name", "tom", nil)
		tx.Set("user:1:name", "Randi", nil)
		tx.Set("user:2:name", "jane", nil)
		tx.Set("user:4:name", "Janet", nil)
		tx.Set("user:5:name", "Paula", nil)
		tx.Set("user:6:name", "peter", nil)
		tx.Set("user:7:name", "Terri", nil)
		return nil
	})
	if err != nil {
		return js.ValueOf(err.Error())
	}

	return js.ValueOf("")
}
