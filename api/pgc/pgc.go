package pg

import (
	"context"
	"fmt"
	"log"
	"os"

	t "github.com/Metzark/cfb/api/types"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Postgres client
type pgc struct {
	db *pgxpool.Pool
}

// Create a pgc (postgres client) from .env
func CreatePGC() *pgc {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load environment file")
	}

	// Load .env values
	host := os.Getenv("CFB_DB_HOST")
    port := os.Getenv("CFB_DB_PORT")
    user := os.Getenv("CFB_DB_USER")
    password := os.Getenv("CFB_DB_PASS")
    name := os.Getenv("CFB_DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

	// Check for missing .env values
	if host == "" || port == "" || user == "" || password == "" || name == "" || sslmode == "" {
		log.Fatal("Failed loading environment values. Not all of the required values are present")
    }

	// Create connection string
	connStr := fmt.Sprintf(
        "postgresql://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, name, sslmode,
    )

	// Create pool
	dbpool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal("Failed creating connection pool")
	}
	
	// Create postgres client
	pg := pgc{ dbpool }

	return &pg
}

// Get a team by their id
func (pg *pgc) GetTeamById(id int) t.Team {
	var team t.Team

	return team
}