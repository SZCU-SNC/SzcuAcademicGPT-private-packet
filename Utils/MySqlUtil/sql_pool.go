package MySqlUtil

import (
	"database/sql"
	"fmt"
	"github.com/SZCU-SNC/SzcuAcademicGPT-private-packet/Utils/ConfigUtil"
	"sync"
	"time"
)

type MySQLPool struct {
	pool           chan *sql.DB
	maxConnections int
	mu             sync.Mutex
}

// InitSQLPool 初始化连接池
func InitSQLPool(maxConnections int) (*MySQLPool, error) {
	pool := make(chan *sql.DB, maxConnections)

	var config = ConfigUtil.GetConfigData()

	var path = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config["database"].(map[interface{}]interface{})["user"],
		config["database"].(map[interface{}]interface{})["password"],
		config["database"].(map[interface{}]interface{})["host"],
		config["database"].(map[interface{}]interface{})["port"],
		config["database"].(map[interface{}]interface{})["dbname"])

	for i := 0; i < maxConnections; i++ {
		// 连接到数据库
		db, err := sql.Open("mysql", path)

		if err != nil {
			return nil, err
		}
		pool <- db
	}

	return &MySQLPool{
		pool:           pool,
		maxConnections: maxConnections,
	}, nil
}

// GetConnectionMysql 获取mysql连接(ps:记得使用ReleaseConnectionMysql释放)
func (p *MySQLPool) GetConnectionMysql() (*sql.DB, error) {
	select {
	case db := <-p.pool:
		return db, nil
	default:
		// Connection pool is empty, check if we can create a new connection
		p.mu.Lock()
		defer p.mu.Unlock()

		if len(p.pool) < p.maxConnections {
			db, err := p.createNewConnection()
			if err != nil {
				return nil, err
			}
			return db, nil
		}

		return nil, fmt.Errorf("connection pool is full")
	}
}

func (p *MySQLPool) ReleaseConnectionMysql(db *sql.DB) {
	select {
	case p.pool <- db:
	default:
		// Connection pool is full, close the extra connection
		db.Close()
	}
}

// ClosePool 关闭连接池中的所有连接
func (p *MySQLPool) ClosePool() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.pool) > 0 {
		db := <-p.pool
		db.Close()
	}
}

// PeriodicallyCloseIdleConnections 定期关闭空闲连接
func (p *MySQLPool) PeriodicallyCloseIdleConnections(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		p.mu.Lock()
		for len(p.pool) > p.maxConnections/2 {
			db := <-p.pool
			db.Close()
		}
		p.mu.Unlock()
	}
}

// createNewConnection 创建新的连接
func (p *MySQLPool) createNewConnection() (*sql.DB, error) {
	var config = ConfigUtil.GetConfigData()

	var path = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config["database"].(map[interface{}]interface{})["user"],
		config["database"].(map[interface{}]interface{})["password"],
		config["database"].(map[interface{}]interface{})["host"],
		config["database"].(map[interface{}]interface{})["port"],
		config["database"].(map[interface{}]interface{})["dbname"])

	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// SetMaxConnections 设置连接池的最大连接数
func (p *MySQLPool) SetMaxConnections(maxConnections int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.maxConnections = maxConnections

	// 增加连接数
	for len(p.pool) < p.maxConnections {
		db, err := p.createNewConnection()
		if err != nil {
			// Handle the error
			fmt.Println("Failed to create new connection:", err)
			return
		}
		p.pool <- db
	}

	// 减少连接数
	for len(p.pool) > p.maxConnections {
		db := <-p.pool
		db.Close()
	}
}
