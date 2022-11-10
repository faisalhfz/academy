package request

type CreateProductRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

type ProductFilterRequest struct {
	Category string `json:"category"`
	MinPrice int    `json:"minprice"`
	MaxPrice int    `json:"maxprice"`
}
