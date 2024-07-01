package product_entity

type Product struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Price int64  `db:"price"`
}
