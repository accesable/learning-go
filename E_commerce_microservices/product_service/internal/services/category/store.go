package category

import (
	"database/sql"
	mysqlc "trann/ecom/product_services/internal/model"
)

type Store struct {
  db *sql.DB
  queries mysqlc.Queries
}

func NewStore(db *sql.DB) *Store {
  return &Store{
    db: db,
    queries: *mysqlc.New(db),
  } 
}



