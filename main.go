package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

var products = []product{
	{ID: "1", Title: "Samsung Galaxy A20s", Price: 20000.00, Description: "Samsung Galaxy A20s black mobile phone"},
	{ID: "2", Title: "Redmi headphone", Price: 1000.00, Description: "Original Redmi headphone pair made in China"},
}

func productById(c *gin.Context) {
	id := c.Param("id")
	book, err := getProductById(id)

	if err != nil {

		c.JSON(404, gin.H{"message": "product not found"})
		return
	}

	c.JSON(200, book)
}

func getProductById(id string) (*product, error) {
	for _, p := range products {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("product not found")
}

func getProducts(c *gin.Context) {
	c.JSON(200, products)
}

func createProduct(c *gin.Context) {
	var newProduct product

	err := c.BindJSON(&newProduct)

	if err != nil {
		return
	}

	// add newProduct to the slice of products
	products = append(products, newProduct)
	c.JSON(200, newProduct)
}

func main() {
	router := gin.Default()
	router.GET("/api", getProducts)
	router.GET("/api/:id", productById)
	router.POST("/api", createProduct)
	router.Run("localhost:3000")

}
