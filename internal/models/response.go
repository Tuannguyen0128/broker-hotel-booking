package models

type Response struct {
	Error *ErrInfo    `json:"error"`
	Body  interface{} `json:"body"`
}

type GetAccountsResp struct {
	Accounts []Account `json:"accounts"`
}
