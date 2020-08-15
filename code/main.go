package main

import (
	"./handler"
	"github.com/labstack/echo"
)

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// ルーティング
	e.GET("/", handler.Index())
	e.GET("/api/hello", handler.Hello())

	// サーバー起動
	e.Start(":8080")
}
