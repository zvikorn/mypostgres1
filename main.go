package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"

	"golang.org/x/sync/errgroup"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "testdb"
)

type customer struct {
	first string    `db:"first_name"`
	last  string    `db:"last_name"`
	time  time.Time `db:"time_created"`
}

func main() {
	fmt.Println("Let's start....")
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set the number of goroutines
	numGoroutines := 3

	// Use an errgroup for synchronization and error handling
	var eg errgroup.Group

	// Use a mutex for safely appending results
	var mu sync.Mutex
	var finalResult []string

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		index := i
		eg.Go(func() error {
			result, err := readDataByIndex(db, index, numGoroutines)
			if err != nil {
				return err
			}

			// Append results using a mutex for safety
			mu.Lock()
			finalResult = append(finalResult, result...)
			mu.Unlock()
			return nil
		})
	}

	// Wait for all goroutines to finish
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	// Process the final result
	for _, data := range finalResult {
		fmt.Println(data)
	}

	//#########Extra - raad one row where customer trans_id = transaction trans_id with join
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal()
	}
	defer tx.Rollback()
	rows, err := tx.Query("SELECT customer.first_name, customer.last_name, transaction.time_created FROM " +
		"customer join transaction on customer.id = transaction.id")
	if err != nil {
		log.Fatal()
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		c := customer{}
		err := rows.Scan(&c.first, &c.last, &c.time)
		if err != nil {
			log.Fatal()
		}
		result = append(result, c.first, c.last, c.time.String())
	}
	for _, r := range result {
		fmt.Println(r)
	}
	//##########################
}

func readDataByIndex(db *sql.DB, index, maxRoutines int) ([]string, error) {
	// Begin a transaction
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query("SELECT first_name, last_name FROM customer WHERE id % $1 = $2", maxRoutines, index)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		c := customer{}
		//var firstName, lastName string
		//err := rows.Scan(&firstName, &lastName)
		err := rows.Scan(&c.first, &c.last)
		if err != nil {
			return nil, err
		}
		//result = append(result, firstName, lastName)
		result = append(result, c.first, c.last)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//
//	_ "github.com/lib/pq" // postgres driver
//)
//
//type customer struct {
//	first string `db:"first_name"`
//	last  string `db:"last_name"`
//	Email string `db:"email"`
//}
//
//func main() {
//	// Connection details
//	host := "127.0.0.1"
//	port := 5432
//	user := "postgres"
//	password := "postgres"
//	dbname := "testdb"
//
//	// Connect to the database
//	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Ping the database to check connection
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Connected to database successfully!")
//
//	// Begin a transaction
//	tx, err := db.Begin()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer tx.Rollback() // Rollback if any error occurs
//
//	// Query the database within the transaction
//	rows, err := tx.Query("SELECT first_name, last_name FROM customer")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	// Process query results
//	for rows.Next() {
//		//var first, last string
//		c := customer{}
//		err := rows.Scan(&c.first, &c.last)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("first: %s, last: %s\n", c.first, c.last)
//	}
//
//	// Check for errors from iterating over rows
//	if err = rows.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//	// Commit the transaction
//	err = tx.Commit()
//	if err != nil {
//		log.Fatal(err)
//	}
//}
