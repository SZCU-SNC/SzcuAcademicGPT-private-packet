package Util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
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

var ENV_YAML string

func GetENV() {
	env, err := GetConfigValue("config.yaml", "environment")
	if err != nil {
		panic(err)
	}
	switch env {
	case "dev":
		ENV_YAML = "config-dev.yaml"
		gin.SetMode(gin.DebugMode)
	case "prod":
		ENV_YAML = "config-prod.yaml"
		gin.SetMode(gin.ReleaseMode)
	case "test":
		ENV_YAML = "config-test.yaml"
		gin.SetMode(gin.TestMode)
	default:
		ENV_YAML = "config.dev.yaml"
		gin.SetMode(gin.DebugMode)
	}
	fmt.Println("ENV_YAML: ", ENV_YAML)
}

func GetConfigValue(filePath, key string) (string, error) {
	// 读取YAML文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// 将YAML内容解析到map中
	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}

	// 如果不包含"."，则直接返回对应的值
	if !strings.Contains(key, ".") {
		return config[key].(string), nil
	}
	// 分解key以支持多级配置
	keys := strings.Split(key, ".")

	// 遍历map来获取对应的一级配置的map
	for i := 0; i < len(keys)-1; i++ {
		config = config[keys[i]].(map[string]interface{})
	}
	// 从一级配置的map中获取对应的值
	return config[keys[len(keys)-1]].(string), nil
}
