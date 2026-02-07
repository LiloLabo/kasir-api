package models

import "time"

type BestSelling struct {
	Name    string `json:"name"`
	QtySold int    `json:"qty_sold"`
}

type Today struct {
	TotalRevenue      int         `json:"total_revenue"`
	TotalTransactions int         `json:"total_transactions"`
	BestSellingItem   BestSelling `json:"best_selling_products"`
}

type ReportData struct {
	ID             int       `json:"id"`
	DateTime       time.Time `json:"datetime"`
	ProductName    string    `json:"product_name"`
	ProductPrice   int       `json:"product_price"`
	Qty            int       `json:"qty"`
	SubTotal       int       `json:"subtotal"`
	RemainingStock int       `json:"remaining_stock"`
}
