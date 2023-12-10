package RedisUtil

/*
配置文件：
redis:
	address: my_addr
	password: my_password
	db: use_default_DB
	poolSize: pool_size
	poolTimeout: 阻塞超时时间 for example: 10s

*/
import (
	"context"
	"fmt"
	"github.com/SZCU-SNC/SzcuAcademicGPT-private-packet/Utils/ConfigUtil"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisClient 结构体中的 pool 字段改为私有
type redisClient struct {
	pool *redis.Client
}

var redisClients redisClient

// InitRedisPool 初始化Redis连接池
func InitRedisPool() error {
	var redisConfig = ConfigUtil.GetConfigData()["redis"].(map[interface{}]interface{})
	var timeOut, _ = time.ParseDuration(redisConfig["poolTimeout"].(string))
	pool := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%v", redisConfig["address"]),
		Password:    fmt.Sprintf("%v", redisConfig["password"]),
		DB:          redisConfig["db"].(int),
		PoolSize:    redisConfig["poolSize"].(int),
		PoolTimeout: timeOut,
	})

	// 测试连接
	pong, err := pool.Ping(context.Background()).Result()
	if err != nil {
		return fmt.Errorf("连接Redis失败: %s", err)
	}

	fmt.Println("成功连接到Redis:", pong)

	// 将 Redis 客户端存储在 redisClients 中
	redisClients = redisClient{pool: pool}

	return nil
}

// Set 设置键值对
func Set(key, value string, exp time.Duration) error {
	err := redisClients.pool.Set(context.Background(), key, value, exp).Err()
	if err != nil {
		return fmt.Errorf("设置键值对失败: %s", err)
	}
	return nil
}

// Get 获取键的值
func Get(key string) (string, error) {
	val, err := redisClients.pool.Get(context.Background(), key).Result()
	if err != nil {
		return "", fmt.Errorf("获取键值失败: %s", err)
	}
	return val, nil
}

// Del 删除键
func Del(key string) error {
	err := redisClients.pool.Del(context.Background(), key).Err()
	if err != nil {
		return fmt.Errorf("删除键失败: %s", err)
	}
	return nil
}

// HSet 设置哈希表字段的值
func HSet(key, field, value string) error {
	err := redisClients.pool.HSet(context.Background(), key, field, value).Err()
	if err != nil {
		return fmt.Errorf("设置哈希表字段的值失败: %s", err)
	}
	return nil
}

// HGet 获取哈希表字段的值
func HGet(key, field string) (string, error) {
	val, err := redisClients.pool.HGet(context.Background(), key, field).Result()
	if err != nil {
		return "", fmt.Errorf("获取哈希表字段的值失败: %s", err)
	}
	return val, nil
}

// HGetAll 获取哈希表的所有字段和值
func HGetAll(key string) (map[string]string, error) {
	val, err := redisClients.pool.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, fmt.Errorf("获取哈希表的所有字段和值失败: %s", err)
	}
	return val, nil
}

// HDel 删除哈希表字段
func HDel(key string, fields ...string) error {
	err := redisClients.pool.HDel(context.Background(), key, fields...).Err()
	if err != nil {
		return fmt.Errorf("删除哈希表字段失败: %s", err)
	}
	return nil
}

// LPush 将值推入列表的左端
func LPush(key string, values ...string) error {
	err := redisClients.pool.LPush(context.Background(), key, values).Err()
	if err != nil {
		return fmt.Errorf("将值推入列表的左端失败: %s", err)
	}
	return nil
}

// RPush 将值推入列表的右端
func RPush(key string, values ...string) error {
	err := redisClients.pool.RPush(context.Background(), key, values).Err()
	if err != nil {
		return fmt.Errorf("将值推入列表的右端失败: %s", err)
	}
	return nil
}

// LRange 获取列表的值
func LRange(key string, start, stop int64) ([]string, error) {
	vars, err := redisClients.pool.LRange(context.Background(), key, start, stop).Result()
	if err != nil {
		return nil, fmt.Errorf("获取列表的值失败: %s", err)
	}
	return vars, nil
}

// LLen 获取列表的长度
func LLen(key string) (int64, error) {
	length, err := redisClients.pool.LLen(context.Background(), key).Result()
	if err != nil {
		return 0, fmt.Errorf("获取列表的长度失败: %s", err)
	}
	return length, nil
}

// LRem 删除列表中值为 value 的元素 count 次
func LRem(key string, count int64, value string) (int64, error) {
	removedCount, err := redisClients.pool.LRem(context.Background(), key, count, value).Result()
	if err != nil {
		return 0, fmt.Errorf("删除列表中值为 %s 的元素失败: %s", value, err)
	}
	return removedCount, nil
}

// LPop 移除并返回列表 key 的头元素
func LPop(key string) (string, error) {
	element, err := redisClients.pool.LPop(context.Background(), key).Result()
	if err != nil {
		return "", fmt.Errorf("移除并返回列表 %s 的头元素失败: %s", key, err)
	}
	return element, nil
}

// RPop 移除并返回列表 key 的尾元素
func RPop(key string) (string, error) {
	element, err := redisClients.pool.RPop(context.Background(), key).Result()
	if err != nil {
		return "", fmt.Errorf("移除并返回列表 %s 的尾元素失败: %s", key, err)
	}
	return element, nil
}

// LTrim 保留列表中从下标 start 到 stop 之间的元素
func LTrim(key string, start, stop int64) error {
	err := redisClients.pool.LTrim(context.Background(), key, start, stop).Err()
	if err != nil {
		return fmt.Errorf("保留列表中从下标 %d 到 %d 之间的元素失败: %s", start, stop, err)
	}
	return nil
}

// SAdd 向集合添加一个或多个成员
func SAdd(key string, members ...string) error {
	err := redisClients.pool.SAdd(context.Background(), key, members).Err()
	if err != nil {
		return fmt.Errorf("向集合添加成员失败: %s", err)
	}
	return nil
}

// SRem 从集合中移除一个或多个成员
func SRem(key string, members ...string) error {
	err := redisClients.pool.SRem(context.Background(), key, members).Err()
	if err != nil {
		return fmt.Errorf("从集合中移除成员失败: %s", err)
	}
	return nil
}

// SMembers 获取集合的所有成员
func SMembers(key string) ([]string, error) {
	members, err := redisClients.pool.SMembers(context.Background(), key).Result()
	if err != nil {
		return nil, fmt.Errorf("获取集合成员失败: %s", err)
	}
	return members, nil
}
