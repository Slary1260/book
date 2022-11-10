/*
 * @Author: tj
 * @Date: 2022-11-10 14:59:28
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 16:44:58
 * @FilePath: \book\database\impl\read.go
 */
package impl

import (
	"fmt"
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

		fmt.Println("Order by age")
		datail = make([]string, 0)
		tx.Ascend("age", func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)

		fmt.Println("Order by age range 30-50")
		datail = make([]string, 0)
		tx.AscendRange("age", `{"age":30}`, `{"age":50}`, func(key, value string) bool {
			fmt.Printf("%s: %s\n", key, value)
			datail = append(datail, value)
			return true
		})
		ret = append(ret, datail...)

		return nil
	})

	return js.ValueOf(strings.Join(ret, ","))
}
