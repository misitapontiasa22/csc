package model

type Product struct {
	Id    int64
	Name  string
	Price int64
	Qty   int64
}

type ProductResponse struct {
	Code    int
	Message string
	Product Product
}
