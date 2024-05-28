package main

import (
	"context"
	"database/sql"
	"e_commerce/product_services/internal/db/product"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// router := gin.Default()

	// router.Run("localhost:8080")
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:my-secret-pw@/e_commerces_microservices_by_go?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	queries := product.New(db)

	// list all authors
	authors, err := queries.ListProducts(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(authors)
}
