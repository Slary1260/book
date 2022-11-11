/*
 * @Author: tj
 * @Date: 2022-11-10 14:41:14
 * @LastEditors: tj
 * @LastEditTime: 2022-11-11 15:31:04
 * @FilePath: \book\database\impl\impl.go
 */
package impl

import (
	"fmt"
	"sync"

	"github.com/tidwall/buntdb"
)

var (
	instance *MemoryDb
	once     sync.Once
)

type MemoryDb struct {
	db *buntdb.DB

	mutex sync.RWMutex
	// 没有创建索引是查询不到数据的
	indexMap   map[string]func(a, b string) bool // 索引key -> 索引函数
	patternMap map[string]string                 // 索引key->索引模型 pattern

	lessMutex    sync.RWMutex
	indexLessMap map[int]func(a, b string) bool // seq -> 索引函数 0:IndexJSON
}

func GetMemoryDb() (*MemoryDb, error) {
	var err error
	once.Do(func() {
		db, e := buntdb.Open(":memory:")
		if e != nil {
			err = e
			return
		}

		instance = &MemoryDb{
			db:           db,
			mutex:        sync.RWMutex{},
			indexMap:     make(map[string]func(a string, b string) bool, 4),
			patternMap:   map[string]string{},
			lessMutex:    sync.RWMutex{},
			indexLessMap: make(map[int]func(a string, b string) bool, 4),
		}
	})
	if err != nil {
		fmt.Println("GetMemoryDb error:", err)
		return nil, err
	}

	return instance, nil
}

func (m *MemoryDb) getIndexLess() {
	m.lessMutex.Lock()
	defer m.lessMutex.Unlock()

	// 0:IndexJSON
	m.indexLessMap[1] = buntdb.IndexBinary
	m.indexLessMap[2] = buntdb.IndexFloat
	m.indexLessMap[3] = buntdb.IndexInt
	m.indexLessMap[4] = buntdb.IndexString
	m.indexLessMap[5] = buntdb.IndexUint
}
