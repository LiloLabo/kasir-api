package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReport() (*models.Today, error) {
	return s.repo.GetReport()
}

func (s *ReportService) GetReportDate(start_date string, end_date string) ([]models.ReportData, error) {
	return s.repo.GetReportDate(start_date, end_date)
}
