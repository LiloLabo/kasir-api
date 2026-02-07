package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetReport() (*models.Today, error) {
	query :=
		`
			select sum(p.subtotal) as revenue, count(id) as total_transactions
			from transaction_details p
		`

	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRevenue := 0
	totalTransactions := 0
	for rows.Next() {
		var revenue, transaction int
		err := rows.Scan(&revenue, &transaction)
		if err != nil {
			return nil, err
		}
		totalRevenue += revenue
		totalTransactions += transaction
	}

	var BestSelling models.BestSelling
	queryBestSelling :=
		`
			select max(p.quantity) as qty_sold, pd.name
			from transaction_details p
			left join products pd on pd.id = p.product_id
			group by pd.name
			ORDER BY qty_sold DESC
			limit 1
		`
	rowsBestSellling, err := repo.db.Query(queryBestSelling)
	if err != nil {
		return nil, err
	}
	defer rowsBestSellling.Close()
	for rowsBestSellling.Next() {
		var qtySold int
		var itemName string
		err := rowsBestSellling.Scan(&qtySold, &itemName)
		if err != nil {
			return nil, err
		}

		BestSelling = models.BestSelling{
			Name:    itemName,
			QtySold: qtySold,
		}
	}

	return &models.Today{
		TotalRevenue:      totalRevenue,
		TotalTransactions: totalTransactions,
		BestSellingItem:   BestSelling,
	}, nil
}

func (repo *ReportRepository) GetReportDate(start_date string, end_date string) ([]models.ReportData, error) {
	query :=
		`
			select p.id, t.created_at as datetime, pd.name, pd.price, p.quantity, p.subtotal, pd.stock
			from transaction_details p
			left join transactions t on t.id = p.transaction_id
      		left join products pd on pd.id = p.product_id
		`
	args := []interface{}{}
	if start_date != "" && end_date != "" {
		query += " WHERE t.created_at >= $1 and t.created_at <= $2"
		args = append(args, start_date, end_date)
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	datareport := make([]models.ReportData, 0)
	for rows.Next() {
		var p models.ReportData
		err := rows.Scan(&p.ID, &p.DateTime, &p.ProductName, &p.ProductPrice, &p.Qty, &p.SubTotal, &p.RemainingStock)
		if err != nil {
			return nil, err
		}
		datareport = append(datareport, p)
	}

	return datareport, nil
}
