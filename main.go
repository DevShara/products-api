package main

import (
	"net/http"
	"products-api/controller"
	"products-api/model"

	"github.com/gin-gonic/gin"
)


func getProducts(c *gin.Context) {
	allProducts := controller.GetAllProducts()
	c.JSON(200, allProducts)
}

func createProduct(c *gin.Context) {
	var newProduct model.Product
	err := c.BindJSON(&newProduct)

	if err != nil {
		return
	}
	controller.InsertProduct(newProduct)
}

func updateProductPrice(c *gin.Context) {

	var input model.Product

	err := c.ShouldBindJSON(&input);
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	  }
	  
	//priceInFloat, _ := strconv.ParseFloat(price, 64)
	controller.UpdateProductPrice(input)
}


func DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	controller.DeleteProduct(id)
}




func main() {
	//connect to MongoDB 
	controller.Init()

	//Create a router instance
	router := gin.Default()

	//Get all products
	router.GET("/api", getProducts)

	//Add a product
	router.POST("/api", createProduct)

	//Update the product price
	router.PUT("/api/", updateProductPrice)

	//Delete a product
	router.DELETE("/api/:id", DeleteProduct)

	//Listening and serving HTTP request
	router.Run("localhost:3000")

}
