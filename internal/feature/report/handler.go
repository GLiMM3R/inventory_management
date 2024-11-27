package report

import (
	"inverntory_management/internal/types"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type ReportHandler struct {
	service ReportServiceImpl
}

func NewSaleHandler(service ReportServiceImpl) *ReportHandler {
	return &ReportHandler{service: service}
}

func (handler *ReportHandler) GetSalesReport(c echo.Context) error {
	var err error

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || page <= 0 {
		limit = 10
	}

	startDateStr := c.QueryParam("start_date")
	var startDate time.Time

	if startDateStr == "" {
		startDate = time.Now()
	} else {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid start date format"})
		}
	}
	startDateUnix := startDate.Unix()

	endDateStr := c.QueryParam("end_date")
	var endDate time.Time

	if endDateStr == "" {
		endDate = time.Now()
	} else {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid end date format"})
		}
	}
	endDateUnix := endDate.Unix()

	response, total, err := handler.service.SalesReport(startDateUnix, endDateUnix, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get sales report"})
	}
	return c.JSON(http.StatusOK, types.Response{
		Data:     response,
		Status:   http.StatusOK,
		Messages: "Success",
		Total:    &total,
	})
}
