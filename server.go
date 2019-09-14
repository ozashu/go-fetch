package main

import (
	"github.com/labstack/echo"
	"github.com/ozashu/go-fetch/router"
)

//main関数を定義しています。プログラムの実行時には、この関数が呼び出されます。
func main() {
	// Echo のインスタンスを作成
	e := echo.New()
	// routerパッケージの Bind関数でルーティングの設定を行っている
	router.Bind(e)
	// publicディレクトリ内のファイルを静的リソースとして配置
	e.Static("/", "public")
	// 1323ポートでサーバーを起動
	e.Logger.Fatal(e.Start(":1323"))
}
