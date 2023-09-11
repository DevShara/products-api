package main

import "github.com/gin-gonic/gin"

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

func getProducts(c *gin.Context) {
	c.JSON(200, products)
}

func main() {
	router := gin.Default()
	router.GET("/api", getProducts)
	router.Run("localhost:3000")

}
