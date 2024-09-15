package report

type SalesReport struct {
	OrderNumber   string  `json:"order_number"`
	TotalQuantity float64 `json:"total_quantity"`
	NetAmount     float64 `json:"net_amount"`
	SaleDate      int64   `json:"sale_date"`
}
