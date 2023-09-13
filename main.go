package main

import (
	"errors"
	"net/http"
	"products-api/controller"
	"products-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var products = []model.Product{
// 	{ID: "1", Title: "Samsung Galaxy A20s", Price: 20000.00, Description: "Samsung Galaxy A20s black mobile phone"},
// 	{ID: "2", Title: "Redmi headphone", Price: 1000.00, Description: "Original Redmi headphone pair made in China"},
// }

func productById(c *gin.Context) {
	id := c.Param("id")
	product, err := getProductById(id)

	if err != nil {

		c.JSON(404, gin.H{"message": "product not found"})
		return
	}

	c.JSON(200, product)
}

func getProductById(id string) (*model.Product, error) {
	// for _, p := range products {
	// 	if p.ID == id {
	// 		return &p, nil
	// 	}
	// }

	return nil, errors.New("product not found")
}

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

	// add newProduct to the slice of products
	// products = append(products, newProduct)
	// c.JSON(200, newProduct)

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

	// if err != ä
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product not found"})
	// 	return
	// }	
}

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	controller.DeleteProduct(id)

	// for i, p := range products {
	// 	if p.ID == id {
	// 		products[i] = products[len(products)-1]     // Copy last element to index i.
	// 		products[len(products)-1] = model.Product{} // Erase last element (write zero value).
	// 		products = products[:len(products)-1]
	// 	}
	// }
}

func main() {

	controller.Init()

	router := gin.Default()
	router.GET("/api", getProducts)
	router.GET("/api/:id", productById)
	router.POST("/api", createProduct)
	router.PUT("/api/:id", updateProductPrice)
	router.DELETE("/api/:id", DeleteProduct)
	router.Run("localhost:3000")

}
