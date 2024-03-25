package product_postgres

const (
	createProduct     = `INSERT INTO products (id, name, price) VALUES ($1, $2, $3)`
	getProductByID    = `SELECT * FROM products WHERE id = $1`
	getAllProduct     = `SELECT * FROM products`
	updateProductByID = `UPDATE products set name = $1, price = $1`
	deleteProductByID = `DELETE FROM products where id = $1`
)
