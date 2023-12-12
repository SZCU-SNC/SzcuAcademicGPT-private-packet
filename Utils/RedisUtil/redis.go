// // Package RedisUtil
// /*
// 配置文件：
// redis:
//
//	address: my_addr
//	password: my_password
//	db: use_default_DB
//
// */
package RedisUtil

//
//import (
//	"context"
//	"fmt"
//	"github.com/SZCU-SNC/SzcuAcademicGPT-private-packet/Utils/ConfigUtil"
//	"github.com/redis/go-redis/v9"
//	"time"
//)
//
//type RedisClient struct {
//	client *redis.Client
//}
//
//// NewRedisClient 创建Redis客户端
//func NewRedisClient() (*RedisClient, error) {
//	var redisConfig = ConfigUtil.GetConfigData()["redis"].(map[interface{}]interface{})
//	client := redis.NewClient(&redis.Options{
//		Addr:     fmt.Sprintf("%v", redisConfig["address"]),
//		Password: fmt.Sprintf("%v", redisConfig["password"]),
//		DB:       redisConfig["db"].(int),
//	})
//
//	// 测试连接
//	pong, err := client.Ping(context.Background()).Result()
//	if err != nil {
//		return nil, fmt.Errorf("连接Redis失败: %s", err)
//	}
//
//	fmt.Println("成功连接到Redis:", pong)
//
//	return &RedisClient{
//		client: client,
//	}, nil
//}
//
//// Close 关闭Redis客户端连接
//func (rc *RedisClient) Close() error {
//	return rc.client.Close()
//}
//
//// Set 设置键值对
//func (rc *RedisClient) Set(key, value string) error {
//	err := rc.client.Set(context.Background(), key, value, 0).Err()
//	if err != nil {
//		return fmt.Errorf("设置键值对失败: %s", err)
//	}
//	return nil
//}
//
//// SetWihExpiration 设置有到期时间的键值对
//func (rc *RedisClient) SetWihExpiration(key, value string, exp time.Duration) error {
//	err := rc.client.Set(context.Background(), key, value, exp).Err()
//	if err != nil {
//		return fmt.Errorf("设置键值对失败: %s", err)
//	}
//	return nil
//}
//
//// Get 获取键的值
//func (rc *RedisClient) Get(key string) (string, error) {
//	val, err := rc.client.Get(context.Background(), key).Result()
//	if err != nil {
//		return "", fmt.Errorf("获取键值失败: %s", err)
//	}
//	return val, nil
//}
//
//// Del 删除键
//func (rc *RedisClient) Del(key string) error {
//	err := rc.client.Del(context.Background(), key).Err()
//	if err != nil {
//		return fmt.Errorf("删除键失败: %s", err)
//	}
//	return nil
//}
//
//// HSet 设置哈希表字段的值
//func (rc *RedisClient) HSet(key, field, value string) error {
//	err := rc.client.HSet(context.Background(), key, field, value).Err()
//	if err != nil {
//		return fmt.Errorf("设置哈希表字段的值失败: %s", err)
//	}
//	return nil
//}
//
//// HGet 获取哈希表字段的值
//func (rc *RedisClient) HGet(key, field string) (string, error) {
//	val, err := rc.client.HGet(context.Background(), key, field).Result()
//	if err != nil {
//		return "", fmt.Errorf("获取哈希表字段的值失败: %s", err)
//	}
//	return val, nil
//}
//
//// LPush 将值推入列表的左端
//func (rc *RedisClient) LPush(key string, values ...string) error {
//	err := rc.client.LPush(context.Background(), key, values).Err()
//	if err != nil {
//		return fmt.Errorf("将值推入列表的左端失败: %s", err)
//	}
//	return nil
//}
//
//// RPush 将值推入列表的右端
//func (rc *RedisClient) RPush(key string, values ...string) error {
//	err := rc.client.RPush(context.Background(), key, values).Err()
//	if err != nil {
//		return fmt.Errorf("将值推入列表的右端失败: %s", err)
//	}
//	return nil
//}
//
//// LRange 获取列表的值
//func (rc *RedisClient) LRange(key string, start, stop int64) ([]string, error) {
//	vars, err := rc.client.LRange(context.Background(), key, start, stop).Result()
//	if err != nil {
//		return nil, fmt.Errorf("获取列表的值失败: %s", err)
//	}
//	return vars, nil
//}
//
//// LLen 获取列表的长度
//func (rc *RedisClient) LLen(key string) (int64, error) {
//	length, err := rc.client.LLen(context.Background(), key).Result()
//	if err != nil {
//		return 0, fmt.Errorf("获取列表的长度失败: %s", err)
//	}
//	return length, nil
//}
