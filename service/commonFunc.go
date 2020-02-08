package service

import (
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"../entity"
)

// Jsonファイルの読み込み＆解析
func readJsonFile() *simplejson.Json {
	raw, err := ioutil.ReadFile("")
	if err != nil {
		fmt.Println(err.Error())
	}

	json, err := simplejson.NewJson(raw)

	return json
}

// 指定したパラメータでJosnを解析して返却する
// 検索区分1でアプリ名、2でユーザ名で検索をする
func searchInJson(param string, searchDiv string) []entity.SearchResult {

	// jsonの読み込み
	raw, err := ioutil.ReadFile("./account.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 構造体へマッピング
	var jsonData []entity.SearchResult
	json.Unmarshal(raw, &jsonData)

	// 結果返却用の構造体を定義
	var result []entity.SearchResult

	// アプリ名で検索
	if searchDiv == "1" {
		for _, stru := range jsonData {
			if stru.AppName == param {
				result = append(result, stru)
			}
		}
		return result
		// ユーザ名で検索
	} else if searchDiv == "2" {
		for _, stru := range jsonData {
			if stru.UserId == param {
				result = append(result, stru)
			}
		}
		return result
	}

	// 想定外の検索区分の場合
	return jsonData
}