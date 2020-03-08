package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)

type Config struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

var dbs sync.Map
var dbLock sync.Mutex

func Init(name string, config Config) error {
	load, ok := dbs.Load(name)
	if ok && load != nil {
		return nil
	}

	dbLock.Lock()
	defer dbLock.Unlock()
	connect, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", config.User, config.Pass, config.Host, config.Port, config.Name))
	if err != nil {
		return err
	}
	dbs.Store(name, connect)
	return nil
}

func DB(name string) *sqlx.DB {
	load, ok := dbs.Load(name)
	if ok && load != nil {
		db, ok := load.(*sqlx.DB)
		if ok {
			return db
		}

	}
	panic(fmt.Sprintf("db [%s] is not exist or not init", name))
}

func Close(name string) error {
	dbLock.Lock()
	defer dbLock.Unlock()
	load, ok := dbs.Load(name)
	if ok && load != nil {
		dbs.Delete(name)
		db, ok := load.(*sqlx.DB)
		if ok {
			return db.Close()
		}
	}
	return nil
}
