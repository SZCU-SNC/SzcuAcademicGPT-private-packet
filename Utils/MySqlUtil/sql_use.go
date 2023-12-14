package MySqlUtil

import (
	"database/sql"
)

// QueryRow 执行单行查询
func QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	// 获取连接
	db, err := mySqlPool.GetConnectionMysql()
	if err != nil {
		return nil, err
	}
	// 释放连接
	defer mySqlPool.ReleaseConnectionMysql(db)

	row := db.QueryRow(query, args...)

	if row.Err() != nil {
		return nil, err
	}

	return row, nil
}

// Insert 执行插入操作
func Insert(query string, args ...interface{}) (int64, error) {
	db, err := mySqlPool.GetConnectionMysql()
	if err != nil {
		return 0, err
	}
	defer mySqlPool.ReleaseConnectionMysql(db)

	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	// 获取插入后的自增ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update 执行更新操作
func Update(query string, args ...interface{}) (int64, error) {
	db, err := mySqlPool.GetConnectionMysql()
	if err != nil {
		return 0, err
	}
	defer mySqlPool.ReleaseConnectionMysql(db)

	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// BulkInsert 执行批量插入操作
func BulkInsert(queries []string, args [][]interface{}) error {
	db, err := mySqlPool.GetConnectionMysql()
	if err != nil {
		return err
	}
	defer mySqlPool.ReleaseConnectionMysql(db)

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for i, query := range queries {
		_, err := tx.Exec(query, args[i]...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
