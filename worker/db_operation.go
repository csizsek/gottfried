package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/csizsek/gottfried/common"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DBOperation struct {}

func (o *DBOperation) DBStore(arg *common.DBStoreArg, reply *common.DBStoreResult) error {
	log.Printf("DBStore data: %s", arg.Data)
	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", CFG.DB.User, CFG.DB.Password, CFG.DB.Host, CFG.DB.Port, CFG.DB.Database)
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		return errors.New("Unable to list bucket\n" + err.Error())
	}
	defer db.Close()
	_, err = db.Exec(fmt.Sprintf("insert into %s values ('%s')", CFG.DB.Table, arg.Data))
	if err != nil {
		return errors.New("Unable to list bucket\n" + err.Error())
	}
	return nil
}
