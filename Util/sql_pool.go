package Util

import (
	"database/sql"
	"fmt"
	"sync"
)

type MySQLPool struct {
	pool     chan *sql.DB
	maxConns int
	mu       sync.Mutex
}

func InitSQLPool(username, password, host, port, database string, maxConns int) (*MySQLPool, error) {
	pool := make(chan *sql.DB, maxConns)
	for i := 0; i < maxConns; i++ {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database))
		if err != nil {
			return nil, err
		}
		pool <- db
	}

	return &MySQLPool{
		pool:     pool,
		maxConns: maxConns,
	}, nil
}

func (p *MySQLPool) GetConnection() (*sql.DB, error) {
	db := <-p.pool
	return db, nil
}

func (p *MySQLPool) ReleaseConnection(db *sql.DB) {
	p.pool <- db
}
