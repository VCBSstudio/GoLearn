package order

type Address struct {
    Street  string `json:"street" binding:"required"`
    City    string `json:"city" binding:"required"`
    ZipCode string `json:"zip_code" binding:"required"`
}

type CreateOrderRequest struct {
    UserID      int     `json:"user_id" binding:"required"`
    ProductID   int     `json:"product_id" binding:"required"`
    Quantity    int     `json:"quantity" binding:"required,gt=0"`
    ShipAddress Address `json:"ship_address" binding:"required"`
}