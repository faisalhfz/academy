package response

type GetCartResponse struct {
	Products   []GetProductResponse `json:"products"`
	TotalPrice int                  `json:"totalprice"`
}
