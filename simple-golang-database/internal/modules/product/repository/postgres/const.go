package product_postgres

const (
	createProduct     = `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, name, price`
	getProductByID    = `SELECT * FROM products WHERE id = $1`
	getAllProduct     = `SELECT * FROM products`
	updateProductByID = `UPDATE products set name = $2, price = $3 WHERE id = $1`
	deleteProductByID = `DELETE FROM products where id = $1`
)
