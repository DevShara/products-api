package main

import (
	"net/http"
	"products-api/controller"
	"products-api/model"
	"strconv"

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
	id := c.Param("id")
	price, priceOk := c.GetQuery("price")

	//TODO - add multiple update 
	//title, titleOk := c.GetQuery("title")

	// if titleOk == false {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing title query param"})
	// 	return
	// }

	if priceOk == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param"})
		return
	}

	priceInFloat, _ := strconv.ParseFloat(price, 64)
	controller.UpdateProductPrice(id, priceInFloat)
}


func DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	controller.DeleteProduct(id)
}

func productById(c *gin.Context) {
	//TODO
	// id := c.Param("id")
	// product, err := getProductById(id)

	// if err != nil {
	// 	c.JSON(404, gin.H{"message": "product not found"})
	// 	return
	// }

	// c.JSON(200, product)
}

// func getProductById(id string) (*model.Product, error) {
// 	

// 	return nil, errors.New("product not found")
// }


func main() {
	//connect to MongoDB 
	controller.Init()

	//Create a router instance
	router := gin.Default()

	//Get all products
	router.GET("/api", getProducts)

	//Get a single product
	router.GET("/api/:id", productById)

	//Add a product
	router.POST("/api", createProduct)

	//Update the product price
	router.PUT("/api/:id", updateProductPrice)

	//Delete a product
	router.DELETE("/api/:id", DeleteProduct)

	//Listening and serving HTTP request
	router.Run("localhost:3000")

}
