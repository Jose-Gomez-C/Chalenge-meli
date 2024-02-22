package model

type BodyItem struct {
	Price      float64 `json:"price"`
	CategoryID string  `json:"category_id"`
	CurrencyID string  `json:"currency_id"`
	SellerID   int     `json:"seller_id"`
	Id         string  `json:"id"`
}

type ResponseItem struct {
	Code int      `json:"code"`
	Body BodyItem `json:"body"`
}

type ResponseCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ResponseCurrency struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type ResponseUser struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
}
