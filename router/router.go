package router

import (
	"github.com/labstack/echo"
	"github.com/ozashu/go-fetch/handler/api"
}

// Bind関数を定義。URL に対応するハンドラーをバインドする。
func Bind(e *echo.Echo) {
	// apiGroupというグループを作成。グループは URL ゙言うところのパスに相当するもので、apiというパスを定義
	apiGroup := e.Group("api")
	// handler/apiパッケージの BindFetchHandler関数でハンドラーの割り当て
	api.BindFetchHandler(apiGroup)
}
