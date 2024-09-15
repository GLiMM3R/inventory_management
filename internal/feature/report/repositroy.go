package report

import (
	"gorm.io/gorm"
)

type ReportRepositoryImpl interface {
	SalesReport(branchID string, startDate, endDate int64, page, limit int) ([]SalesReport, int64, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepositoryImpl {
	return &reportRepository{db: db}
}

// SalesReport implements IReportRepositoryImpl.
func (r *reportRepository) SalesReport(branchID string, startDate, endDate int64, page, limit int) ([]SalesReport, int64, error) {
	var salesReports []SalesReport
	var total int64
	offset := (page - 1) * limit

	query := r.db.Table("sales")

	query = query.Select("sales.order_number as order_number, SUM(sales.total_price) as net_amount", "SUM(sales.quantity) as total_quantity", "sales.sale_date as sale_date").
		Joins("left join inventories on inventories.inventory_id = sales.fk_inventory_id").
		Where("inventories.fk_branch_id = ? AND sales.sale_date >= ? AND sales.sale_date <= ?", branchID, startDate, endDate).
		Group("order_number, sale_date").
		Count(&total).Limit(limit).Offset(offset).
		Scan(&salesReports)

	if err := query.Error; err != nil {
		return nil, 0, err
	}

	return salesReports, total, nil
}
