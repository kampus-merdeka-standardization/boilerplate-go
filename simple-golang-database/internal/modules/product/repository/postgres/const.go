package product_postgres

const (
	createProduct     = `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, name, price`
	getProductByID    = `SELECT id, name, price FROM products WHERE id = $1`
	getAllProduct     = `SELECT id, name, price FROM products`
	updateProductByID = `UPDATE products set name = $2, price = $3 WHERE id = $1 RETURNING id, name, price`
	deleteProductByID = `DELETE FROM products where id = $1`
)
