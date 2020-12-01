package repository

import "service/models"

var products []models.Product
var nextID = 1

// GetProducts is for Getting a product
func GetProducts() []models.Product {
	return products
}

// AddProduct is for Adding product
func AddProduct(product models.Product) int {
	product.ID = nextID
	nextID++
	products = append(products, product)
	return product.ID
}
