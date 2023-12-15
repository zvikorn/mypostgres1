package service

import (
	"context"
	"mypostgres1/pkg/dao"
)

type ScannerService struct {
	ScannerDAO dao.Scanner
}

func NewScannerService() *ScannerService {
	scannerService := &ScannerService{}
	//TODO initialize ScanerService.ScannerDAO
	return scannerService
}

func (ss *ScannerService) createResource() {
	return
}

func (ss *ScannerService) GetURNsByServiceName(ctx context.Context, serviceName string) ([]string, error) {
	return ss.ScannerDAO.GetURNsByServiceName(ctx, serviceName)
}
