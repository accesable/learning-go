// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: product.sql

package mysqlc

import (
	"context"
)

const createProduct = `-- name: CreateProduct :exec
INSERT INTO products (name, description, image, price, quantity)
VALUES (?, ?, ?, ?, ?)
`

type CreateProductParams struct {
	Name        string
	Description string
	Image       string
	Price       string
	Quantity    uint32
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Image,
		arg.Price,
		arg.Quantity,
	)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id uint32) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProductByID = `-- name: GetProductByID :one
SELECT id, name, description, image, price, quantity, createdAt
FROM products
WHERE id = ?
`

func (q *Queries) GetProductByID(ctx context.Context, id uint32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Price,
		&i.Quantity,
		&i.Createdat,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, image, price, quantity, createdAt
FROM products
ORDER BY createdAt DESC
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Image,
			&i.Price,
			&i.Quantity,
			&i.Createdat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET name = ?, description = ?, image = ?, price = ?, quantity = ?
WHERE id = ?
`

type UpdateProductParams struct {
	Name        string
	Description string
	Image       string
	Price       string
	Quantity    uint32
	ID          uint32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Image,
		arg.Price,
		arg.Quantity,
		arg.ID,
	)
	return err
}
