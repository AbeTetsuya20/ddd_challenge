package main

import (
	"context"
	"database/sql"
	"fmt"
	domain "github.com/AbeTetsuya20/ddd_challenge/server/domain/model"
	handler "github.com/AbeTetsuya20/ddd_challenge/server/interface"
	"github.com/go-sql-driver/mysql"
	"net"
	"os"
	"strconv"
)

func main() {
	fmt.Println("run main.go")

	ctx := context.Background()

	// database の初期化
	sql, err := InitDatabase("dockerMySQL", 3306, "app", "root", "password")
	if err != nil {
		fmt.Println("error db Open")
		os.Exit(1)
	}
	defer sql.Close()

	user, _ := sql.QueryContext(ctx, "select * from user")
	for user.Next() {
		var tmp domain.User
		user.Scan(&tmp.UserID, &tmp.UserName, &tmp.Password, &tmp.CreatedAt, &tmp.UpdatedAt)
		fmt.Printf("user: %+v", tmp)
	}

	// interface の初期化
	service := handler.InitService(sql)

	// サーバーの起動
	service.Server(ctx)

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
