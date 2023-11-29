package MySqlUtil

import (
	"database/sql"
	"fmt"
	"github.com/SZCU-SNC/SzcuAcademicGPT-private-packet/Utils/ConfigUtil"
)

func InitializeDatabase() (*sql.DB, error) {

	var config = ConfigUtil.GetConfigData()

	var path = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config["database"].(map[interface{}]interface{})["user"],
		config["database"].(map[interface{}]interface{})["password"],
		config["database"].(map[interface{}]interface{})["host"],
		config["database"].(map[interface{}]interface{})["port"],
		config["database"].(map[interface{}]interface{})["dbname"])
	// 连接到数据库
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return db, nil
}
