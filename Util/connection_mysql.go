package Util

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Environment string `yaml:"environment"`
}

func InitializeDatabase() (*sql.DB, error) {
	envConfigFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("Failed to read the environment config file: %v", err)
	}

	// 解析环境配置
	var envConfig Config
	err = yaml.Unmarshal(envConfigFile, &envConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal the environment config file: %v", err)
	}

	// 选择要加载的配置文件
	var configFile string
	switch envConfig.Environment {
	case "test":
		configFile = "config-test.yaml"
	case "prod":
		configFile = "config-prod.yaml"
	default:
		configFile = "config-dev.yaml"
	}

	// 读取选择的配置文件
	configFileData, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to read the config file: %v", err)
	}

	// 解析选择的配置文件
	var config map[string]interface{}
	err = yaml.Unmarshal(configFileData, &config)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal the config file: %v", err)
	}

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
