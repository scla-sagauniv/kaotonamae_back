package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"size:100;unique"`
}

// DBを起動させる
func gormConnect() (*gorm.DB, error) {
	// DSN (Data Source Name) の詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`,
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB接続に失敗しました:", err)
	} else {
		fmt.Println("DB接続に成功しました:", db)
	}

	return db, nil
}

func main() {
	// // DB起動
	// db := dbInit()

	db, err := gormConnect()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	// マイグレーション
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// fmt.Println("Hello, World!")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/db_test", func(c echo.Context) error {
		// Userテーブル作成
		db.AutoMigrate(&User{})
		return c.String(http.StatusOK, "test")
	})
	e.Logger.Fatal(e.Start(":8080"))

}
