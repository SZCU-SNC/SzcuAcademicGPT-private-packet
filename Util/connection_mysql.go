package Util

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Environment string `yaml:"environment"`
}

func InitializeDatabase() (*sql.DB, error) {

	var config, _ = ReadConfigFromYaml()

	var path = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config["database"].(map[interface{}]interface{})["user"],
		config["database"].(map[interface{}]interface{})["password"],
		config["database"].(map[interface{}]interface{})["host"],
		config["database"].(map[interface{}]interface{})["port"],
		config["database"].(map[interface{}]interface{})["dbname"])
	// 连接到数据库
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the database: %v", err)
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to ping the database: %v", err)
	}

	return db, nil
}
