package main

import (
	"fmt"
	"os"

	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

// MySQLに作成するUserテーブルの定義
type User struct {
	// gorm.Modelをつけると、idとCreatedAtとUpdatedAtとDeletedAtが作られる
	gorm.Model

	Name string
	Age  int
}

// DBを起動させる
func dbInit() *gorm.DB {
	// DSN (Data Source Name) の詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB接続に失敗しました:", err)
	}
	fmt.Println("DB接続に成功しました:", db)

	// 接続に成功した場合、「db connected!!」と表示する
	fmt.Println("db connected!!")
	return db
}

func main() {
	// // DSN (Data Source Name) の詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	// // [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	// dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
	// 	os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// fmt.Println("dsn:", dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("DB接続に失敗しました:", err)
	// 	return
	// }
	// fmt.Println("DB接続に成功しました:", db)

	// DB起動
	db := dbInit()

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
