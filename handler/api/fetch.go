package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ozashu/go-fetch/logic"
)

// fetchRequestという型を定義。fetchRequestはリクエスト情報を保持するための構造体で、string型のURLをメンバーに持つ。
type fetchRequest struct {
	URL string
}

// BindFetchHandler関数を定義。この関数では、URL に対応するハンドラーをバインド
func BindFetchHandler(g *echo.Group) {
	// fetchというパスに対して POST リクエストが送信された場合の処理を定義。fetch関数を実行するようにしている
	g.POST("/fetch", fetch)
}

// fetch関数を定義。この関数では、リクエストを元にウェブページをフェッチし、レスポンスを生成
func fetch(c echo.Context) error {
	// リクエスト情報を fetchRequest型の変数 reqにバインド
	// エラーが発生した 場合は、ステータスコード 400(http.StatusBadRequest)を返す
	req := new(fetchRequest)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// logicパッケージの Fetch関数でウェブページをフェッチ
	// エラーが発生した 場合は、ステータスコード 400(http.StatusBadRequest)を返す
	resp, err := logic.Fetch(req.URL)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// ウェブページのフェッチに成功した場合、ステータスコードは 200(http.StatusOK)で、 取得したレスポンスボディをそのまま返す
	return c.String(http.StatusOK, resp.Body)
}
