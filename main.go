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

// func updateProductPrice(c *gin.Context) {
// 	id, ok := c.GetQuery("id")
// 	price, ok := c.GetQuery("price")

// 	if ok == false {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param"})
// 		return
// 	}

// 	product, err := getProductById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product not found"})
// 		return
// 	}

// 	product.Price = price
// }

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	for i, p := range products {
		if p.ID == id {

			products[i] = products[len(products)-1] // Copy last element to index i.
			products[len(products)-1] = product{}   // Erase last element (write zero value).
			products = products[:len(products)-1]
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/api", getProducts)
	router.GET("/api/:id", productById)
	router.POST("/api", createProduct)
	router.DELETE("/api/:id", DeleteProduct)
	router.Run("localhost:3000")

}
