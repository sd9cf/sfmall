package model

type ProductApiGetProductReq struct {
	ProductId string `json:"productId"`
}

type ProductApiGetProductsReq struct {
	// CategoryId string `json:"categoryId"`
	PagingQueryReq
}

type SimpleProduct struct {
	Id        uint64  `json:"productid"`
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	MainImage string  `json:"mainImage"`
}

type BuyProduct struct {
	ProductIds []uint64 `v:"required#商品不为空"`
	AddressId uint `v:"required#地址不为空"`
}
