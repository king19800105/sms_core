package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strconv"
	"fmt"
)

type DBConfig map[string]string

type MysqlConnection struct {
	Config FileContent
	Engine *xorm.Engine
}

func NewMysqlConnection() DBConnection {
	return &MysqlConnection{}
}

func (connect *MysqlConnection) initDBConnection(config FileContent) DBConnection {
	source := config["username"].(string) + ":" + config["password"].(string) +
		"@tcp(" + config["host"].(string) + ":" + strconv.Itoa(config["port"].(int)) + ")/" +
		config["database"].(string) + "?charset=" + config["charset"].(string)
	connect.Config = config
	engine, err := xorm.NewEngine(config["driver"].(string), source)

	if nil != err {
		fmt.Println(err)
		panic("db connect failed")
	}

	if idle, ok := config["max-idle"].(int); ok {
		engine.SetMaxIdleConns(idle)
	}

	if open, ok := config["max-open"].(int); ok {
		engine.SetMaxOpenConns(open)
	}

	connect.Engine = engine

	return connect
}

