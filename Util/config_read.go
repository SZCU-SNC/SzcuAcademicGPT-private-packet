package Util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// 私有类型，只能在包内访问,防止配置数据被修改
type configData map[string]interface{}
type Config struct {
	Environment string `yaml:"environment"`
}

var config configData

// GetConfigData 获取内存中config
func GetConfigData() configData {
	return config
}

// InitConfig 初始化全局config
func InitConfig() {
	config, _ = ReadConfigFromYaml()
}

// ReadConfigFromYaml 当希望每次都从yaml查询时使用，优点是更新配置文件后不需要重启
func ReadConfigFromYaml() (map[string]interface{}, error) {

	//获取main文件目录
	mainDir, err := os.Getwd()

	envConfigFile, err := os.ReadFile(filepath.Join(mainDir + "config.yaml"))
	if err != nil {
		return nil, fmt.Errorf("failed to read the environment config file: %v", err)
	}

	// 解析环境配置
	var envConfig Config
	err = yaml.Unmarshal(envConfigFile, &envConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the environment config file: %v", err)
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
		return nil, fmt.Errorf("failed to read the config file: %v", err)
	}

	// 解析选择的配置文件
	var config map[string]interface{}
	err = yaml.Unmarshal(configFileData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the config file: %v", err)
	}

	return config, nil
}
