package main

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/nalcheg/beego_app_kickstart/routers"
)

func main() {
	beego.Run()
}

func init() {
	dbFileName := beego.AppConfig.String("dbFileName")

	err := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = orm.RegisterDataBase("default", "sqlite3", "file:"+dbFileName)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if os.Getenv("DEBUG_MODE") == "true" {
		orm.Debug = true
	}

	err = createSchemeIfNotExisted()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func createSchemeIfNotExisted() error {
	o := orm.NewOrm()

	_, err := o.Raw("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, login TEXT, password TEXT)").Exec()
	if err != nil {
		return err
	}

	num, err := o.QueryTable("users").Count()
	if err != nil {
		return err
	}

	if num == 0 {
		_, err = o.Raw(`INSERT INTO users (login, password) VALUES ('odmin','$2a$14$dmCpB2Hync5ghQwScJ/RYe24XjuisBovJDHAPPq9jlkQjpTSydSFi')`).Exec()
	}

	return nil
}
