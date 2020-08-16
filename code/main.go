package main

import (
	// ディレクトリ名を指定
	"./external"
	"github.com/labstack/echo"
)

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// ルーティング
	// ディレクトリ名で呼び出し
	e.GET("/", external.Home())
	e.GET("/hello", external.Hello())

	// サーバー起動
	e.Start(":8080")
}
