package product_entity

type Product struct {
	ID    string  `db:"id"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}