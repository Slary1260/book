/*
 * @Author: tj
 * @Date: 2022-11-10 14:59:28
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 15:31:11
 * @FilePath: \book\database\impl\read.go
 */
package impl

import (
	"strings"
	"syscall/js"

	// jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/buntdb"
)

// var (
// 	json = jsoniter.ConfigCompatibleWithStandardLibrary
// )

func (m *MemoryDb) ReadData(this js.Value, args []js.Value) interface{} {
	ret := make([]string, 0)
	m.db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("names", func(key, val string) bool {
			log.Infof("%s %s\n", key, val)
			ret = append(ret, key)
			return true
		})
		return nil
	})

	return js.ValueOf(strings.Join(ret, ","))
}
