package main

import (
	"fmt"
	"kaotonamae_back/db"
	"kaotonamae_back/migrate"
	"kaotonamae_back/routes"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	colorReset  = "\033[37m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorRed    = "\033[31m"
)

func main() {
	// DB接続
	db.Init()
	migrate.Run()

	server := echo.New()
	routes.RegisterRoutes(server)

	registerMiddleware(server)

	go func() {
		if err := server.Start(":8080"); err != nil {
			server.Logger.Fatal(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	printRoutes(server)
}

func registerMiddleware(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			printRoutes(e)
			return err
		}
	})
}

func printRoutes(e *echo.Echo) {
	fmt.Println("\nRegistered routes:")
	for _, route := range e.Routes() {
		var methodColor string
		switch route.Method {
		case echo.GET:
			methodColor = colorGreen
		case echo.POST:
			methodColor = colorYellow
		case echo.PUT:
			methodColor = colorBlue
		case echo.DELETE:
			methodColor = colorRed
		default:
			methodColor = colorReset
		}

		// ルートのパスを動的パラメータを含む形式に変更する
		path := formatPath(route.Path)

		// メソッドとパスを別々の色で表示
		fmt.Printf("%s%-6s%s %-12s%s\n", methodColor, route.Method, colorReset, path, colorReset)
	}
	fmt.Printf("\n")
}

func formatPath(path string) string {
	// 正規表現で動的な部分を `{}` で囲んだ形式に変換する
	re := regexp.MustCompile(`:([^/]+)`)
	return re.ReplaceAllString(path, "{$1}")
}
