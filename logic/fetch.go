package logic

import (
	"io/ioutil"
	"net/http"
)

// FetchResponseという型を定義
// FetchResponseはレスポンス情報を保持するための構造体で、string型の Bodyをメンバーに持つ
type FetchResponse struct {
	Body string
}

// Fetch関数を定義。引数で受け取った URL を元にウェブページ をフェッチする。
func Fetch(url string) (*FetchResponse, error) {
	// リクエスト情報を持つ変数 reqを作成。エラーが発生した場合は、そのエラー をそのまま返す。
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return nil, err
	}
	// HTTP クライアントを作成
	client := new(http.Client)
	// HTTP クライアントでリクエストを送信。エラーが発生した場合は、そのエラーをそのまま返す。
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// resp.Bodyは ReadCloser型なので、defer 文で Close関数を実行してリソースを解放
	defer resp.Body.Close()
	// ioutilパッケージのReadAll関数でresp.Bodyを読み込んで、[]byte型の変数 bodyに格納
	// エラーが発生した場合は、そのエラーをそのまま返す
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// FetchResponse型の値を作成してそのポインターを返す
	return &FetchResponse{
		Body: string(body),
	}, nil
}
