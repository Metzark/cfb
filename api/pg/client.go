package pg

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Extract actual values (into s2) from a struct with pgtype values (from s1)
func PGExtract[T1 any, T2 any](s1 *T1, s2 *T2) {

	// Values from struct 1
	s1Value := reflect.ValueOf(s1).Elem()

	// Types from struct 1
	s1Type := s1Value.Type()

	// Reflect struct 2 so that the values can be set
	s2Value := reflect.ValueOf(s2).Elem()

	// For each field in struct 1
	for i := 0; i < s1Value.NumField(); i++ {
		// Get field from struct 1
		field := s1Type.Field(i)

		// Get value from struct 1
		value := s1Value.Field(i)

		// Get the corresponding field for struct 2
		s2Field := s2Value.FieldByName(field.Name)

		// Go through pgtypes, check for valid values and set field for struct 2
		switch v := value.Interface().(type) {
		case pgtype.Text:
			if v.Status == pgtype.Present {
				s2Field.Set(reflect.ValueOf(v.String))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Int4:
			if v.Status == pgtype.Present {
				s2Field.Set(reflect.ValueOf(v.Int))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Float8:
			if v.Status == pgtype.Present {
				s2Field.Set(reflect.ValueOf(v.Float))
			} else {
				s2Field.Set(reflect.Zero(s2Field.Type()))
			}
		case pgtype.Bool:
			s2Field.Set(reflect.ValueOf(v.Bool))
		default:
			if s2Field.Type() == value.Type() {
				s2Field.Set(value)
			}
		}
	}
}

// Create a pgc (postgres client) from .env
func CreatePGC() *PGC {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load environment file")
	}

	// Load .env values
	host := os.Getenv("CFB_DB_HOST")
    port := os.Getenv("CFB_DB_PORT")
    user := os.Getenv("CFB_DB_USER")
    password := os.Getenv("CFB_DB_PASSWORD")
    name := os.Getenv("CFB_DB_NAME")
    sslmode := os.Getenv("CFB_DB_SSLMODE")

	// Check for missing .env values
	if host == "" || port == "" || user == "" || password == "" || name == "" || sslmode == "" {
		log.Fatal("Failed loading environment values. Not all of the required values are present.")
    }

	// Create connection string
	connStr := fmt.Sprintf(
        "postgresql://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, name, sslmode,
    )

	// Create pool
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal("Failed creating connection pool")
	}
	
	// Create postgres client
	pg := PGC{ pool }

	return &pg
}

// Close a pgc (postres client)
func (pgc *PGC) ClosePGC() {
	pgc.pool.Close()
}

// Query postgres with the provided client, a parametrized query string, and args
func query[T any](pgc *PGC, q string, args ...interface{}) ([]T, error) {
	// Array of generic type for the rows returned from the query
	var arr []T

	// Query postgres
	rows, err := pgc.pool.Query(context.Background(), q, args...)

	// Return query error if there was one
	if err != nil {
		return arr, fmt.Errorf("error executing query: %v", err)
	}

	// Close the rows before return
	defer rows.Close()

	// Collect rows into the generic type array
	arr, err = pgx.CollectRows(rows, pgx.RowToStructByName[T])

	// Return error if there was one when collecting rows
	if err != nil {
		return arr, fmt.Errorf("error collecting rows after query: %v", err)
	}

	return arr, nil
}

// Query postgres, this function is able to run arbitrary select statements
func customQuery(pgc *PGC, q string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	// Execute query
	rows, err := pgc.pool.Query(context.Background(), q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Get column names from rows
	fields := rows.FieldDescriptions()
	columns := make([]string, len(fields))
	for i, fd := range fields {
		columns[i] = string(fd.Name)
	}

	// Get the values from each row
	for rows.Next() {
		// Get values of the row
		values, err := rows.Values()

		if err != nil {
			return nil, err
		}

		// Create row map
		row := make(map[string]interface{})
		
		// Map out columns and values
		for i, name := range columns {
			row[name] = values[i]
		}

		// Add to results
		results = append(results, row)
	}

	return results, err
}