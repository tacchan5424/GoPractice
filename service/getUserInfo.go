package service

import (
	"../entity"
)

// 指定したアプリで登録しているアカウント情報すべて取得する
func GetUserInfoFromAppName(appName string, searchDiv string) ([]entity.SearchResult, error) {
	userInfo, err := searchInJson(appName, searchDiv)
	return userInfo, err
}

// 指定したユーザIDを使用しているアカウント情報すべて取得する
func GetUserInfoFromUserId(userId string, searchDiv string) ([]entity.SearchResult, error) {
	userInfo, err := searchInJson(userId, searchDiv)
	return userInfo, err
}

// 保有しているアカウント情報すべて取得する
func GetAllUserInfo(searchDiv string) ([]entity.SearchResult, error) {
	userInfo, err := searchInJson("", searchDiv)
	return userInfo, err
}