package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	fmt.Println("run main.go")

	ctx := context.Background()

	// database の初期化
	sql, err := InitDatabase("localhost", 3306, "app", "root", "password")
	if err != nil {
		fmt.Println("error db Open")
		os.Exit(1)
	}

	_, err = sql.ExecContext(ctx, "use app")
	if err != nil {
		log.Fatal(err)
	}

	_, err = sql.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS user (id integer, name varchar(30))")
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO user (id, name) VALUES (?, ?)"
	_, err = sql.ExecContext(ctx, query, 1234, "test")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// interface の初期化

	// サーバーの起動

	fmt.Println("finish main.go")
}

func InitDatabase(host string, port uint16, dbname, username, password string) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.Addr = net.JoinHostPort(host, strconv.Itoa(int(port)))
	cfg.DBName = dbname
	cfg.User = username
	cfg.Passwd = password
	cfg.ParseTime = true

	connector, err := mysql.NewConnector(cfg)
	if err != nil {
		return nil, fmt.Errorf("new connector: %w", err)
	}

	return sql.OpenDB(connector), nil
}
