package main


// Product to tylko model, nie lista produktow
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}


// products bedzie lista
var products = []Product{
	{ID: 1, Name: "Laptop", Price: 3500.0},
	{ID: 2, Name: "Telefon", Price: 2000.0},
	{ID: 3, Name: "Klawiatura", Price: 150.0},
}


