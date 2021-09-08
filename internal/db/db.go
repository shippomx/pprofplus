package db

import (
	"github.com/shippomx/pprofplus/internal/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

type Mem struct {
	types.MemProfile
}

var Cli *gorm.DB

func dbFileExist(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func init() {
	dbFile := "pprof.db"
	Cli, _ = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	Cli.AutoMigrate(&types.MemProfile{})
}
