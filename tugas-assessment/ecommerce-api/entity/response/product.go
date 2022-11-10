package response

type GetProductResponse struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type GetProductDetailResponse struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
