package dao

import (
	"context"
	"time"

	"mypostgres1/pkg/models"
)

type SQLScanner struct{}

type Scanner interface {
	InsertResource(ctx context.Context, spec models.Resource) error
	ListRsourcesByResourceType(ctx context.Context, resourceType int) ([]models.Resource, error)
	GetResourceByURN(ctx context.Context, urn string) (models.Resource, error)
	GetURNsByServiceName(ctx context.Context, serviceName string) ([]string, error)
	UpdateResource(ctx context.Context, urn, resourceType, name string, date time.Time) error
	DeleteResource(ctx context.Context, urn string) error
}
