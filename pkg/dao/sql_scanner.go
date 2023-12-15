package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	"database/sql"

	"mypostgres1/pkg/models"
)

const (
	// Connection details
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "testdb"
)

func (s SQLScanner) InsertResource(ctx context.Context, spec models.Resource) error {
	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	insertQuery := "INSERT INTO resource (urn, type, name, date) VALUES ($1, $2, $3, $4)"

	// Execute the insert query
	_, err = db.Query(insertQuery, spec.URN, spec.ResourceType, spec.Name, spec.Date)
	if err != nil {
		return err
	}
	return nil
}

func (s SQLScanner) ListRsourcesByResourceType(ctx context.Context, resourceType int) ([]models.Resource, error) {
	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return []models.Resource{}, err
	}
	defer db.Close()
	insertQuery := "SELECT * FROM resourcee WHERE resourceType = $1"

	// Execute the insert query
	rows, err := db.Query(insertQuery, resourceType)
	if err != nil {
		return []models.Resource{}, err
	}
	// Process the results
	var resources []models.Resource
	for rows.Next() {
		err := rows.Scan(&resources)
		if err != nil {
			return []models.Resource{}, err
		}
	}
	return resources, nil
}

func (s SQLScanner) GetURNsByServiceName(ctx context.Context, serviceName string) ([]string, error) {

	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return []string{}, err
	}
	defer db.Close()

	// Query the database within the transaction
	rows, err := db.Query("select urns from service where sname=$1", serviceName)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	var urns = []string{}
	for rows.Next() {
		err := rows.Scan(&urns)
		if err != nil {
			log.Fatal(err)
		}
	}
	return urns, nil
}

func (s SQLScanner) UpdateResource(ctx context.Context, urn, resourceType, name string, date time.Time) error {
	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("UPDATE resourcee SET resourceType = $1 where urn=$2", resourceType, urn)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (s SQLScanner) GetResourceByURN(ctx context.Context, urn string) (models.Resource, error) {
	return models.Resource{}, nil
}

func (s SQLScanner) DeleteResource(ctx context.Context, urn string) error { return nil }
