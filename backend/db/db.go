package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dsn := "root:rootpassword@tcp(localhost:3309)/certificate_ledger?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	if err := initTables(db); err != nil {
        return nil, fmt.Errorf("failed to init tables: %v", err)
    }

    return db, nil
}

func initTables(db *sql.DB) error {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id VARCHAR(36) PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
            role VARCHAR(50) NOT NULL DEFAULT 'user',
            created_at DATETIME NOT NULL,
            updated_at DATETIME NOT NULL
        )`,
        `CREATE TABLE IF NOT EXISTS certificates (
            id VARCHAR(50) PRIMARY KEY,
            hash VARCHAR(64) NOT NULL,
            recipient_name VARCHAR(255) NOT NULL,
            recipient_email VARCHAR(255) NOT NULL,
            certificate_title VARCHAR(255) NOT NULL,
            issue_date DATETIME NOT NULL,
            issuer_id VARCHAR(36) NOT NULL,
            issuer_name VARCHAR(255) NOT NULL,
            description TEXT,
            block_number INT NOT NULL,
            timestamp DATETIME NOT NULL,
            FOREIGN KEY (issuer_id) REFERENCES users(id)
        )`,
    }

    for _, query := range queries {
        if _, err := db.Exec(query); err != nil {
            return fmt.Errorf("failed to execute query: %v", err)
        }
    }
    return nil
}