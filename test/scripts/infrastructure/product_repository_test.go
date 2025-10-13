package infrastructure

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang-product-app.git/common/postgresql"
	"golang-product-app.git/persistence"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("Before all test")
	exitcode := m.Run()
	fmt.Println("After all test")
	os.Exit(exitcode)
}

func TestGetAllProducts(t *testing.T) {
	fmt.Println("productRepository")

}
