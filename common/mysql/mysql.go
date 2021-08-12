package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func NewMysql(datasource string) *gorm.DB {
	db, err := gorm.Open("mysql", datasource)
	if err != nil {
		panic(fmt.Sprintf("gorm open error: %s", err))
	}
	return db
}
