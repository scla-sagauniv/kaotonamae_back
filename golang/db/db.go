package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// 環境変数の読み込み
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	// tcp（）の中にdocker-composeで定義したDB用コンテナのサービス名を入れれば、
	// 自動的にホストとポートを読み取ってくれる
	// DSN (Data Source Name) の詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dsn := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local`,
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database.")
	}
}
