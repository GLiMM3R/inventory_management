package report

import "inverntory_management/internal/feature/user"

type ReportServiceImpl interface {
	SalesReport(userID string, startDate, endDate int64, page, limit int) ([]SalesReport, int64, error)
}

type reportService struct {
	repo     ReportRepositoryImpl
	userRepo user.UserRepositoryImpl
}

func NewReportService(repo ReportRepositoryImpl, userRepo user.UserRepositoryImpl) ReportServiceImpl {
	return &reportService{repo: repo, userRepo: userRepo}
}

// SalesReport implements IReportServiceImpl.
func (s *reportService) SalesReport(userID string, startDate, endDate int64, page, limit int) ([]SalesReport, int64, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, 0, err
	}

	reports, total, err := s.repo.SalesReport(user.BranchID, startDate, endDate, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return reports, total, nil
}
