package entity

type SearchParams struct {
	AppName   string  `json:"appName"`
	UserId    string  `json:"userId"`
}

type SearchResult struct {
	AppName   string  `json:"appName"`
	UserId    string  `json:"userId"`
	Password  string  `json:"password"`
}