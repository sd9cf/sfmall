package model

type ProductApiGetProductReq struct {
	CategoryId string `json:"categoryId"`
}

type ProductApiGetProductsReq struct {
	CategoryId string `json:"categoryId"`
	PagingQueryReq
}

type SimpleProduct struct {
	Id        uint64  `json:"id"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
}
