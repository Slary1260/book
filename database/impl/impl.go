/*
 * @Author: tj
 * @Date: 2022-11-10 14:41:14
 * @LastEditors: tj
 * @LastEditTime: 2022-11-10 14:57:08
 * @FilePath: \book\database\impl\impl.go
 */
package impl

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/buntdb"
)

var (
	instance *MemoryDb
	once     sync.Once

	log = logrus.WithFields(logrus.Fields{
		"MemoryDb": "",
	})
)

type MemoryDb struct {
	db *buntdb.DB
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
			db: db,
		}
	})

	if err != nil {
		log.Errorln("GetMemoryDb error:", err)
		return nil, err
	}

	return instance, nil
}
