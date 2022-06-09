package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"	
	"sandbox/api/models"
)

type ProductController struct{}

var productModel = new(models.Product)

// TODO: Add missing layers of bussines and persistance
// TODO: Add model conversion to accept and return models
// TODO: Add error handler layer to create generic controller method with structured API response for success and failure response types
// TODO: Auto generate swagger docs with swag dependency

// @BasePath /v1/products/

// List products
// @Summary List products
// @Schemes
// @Description List the products on th DB
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {string} TDB
// @Router /v1/products/ [get]
func (u ProductController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Products listed!"})
}

// Retrieve a product
// @Summary Retrieve a product
// @Schemes
// @Description Retrieve a product by Id
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {string} TDB
// @Router /v1/products/:productId [get]
func (u ProductController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Product fetched!"})
}

// Create a product
// @Summary Create a product
// @Schemes
// @Description Create a product
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {string} TDB
// @Router /v1/products/ [post]
func (u ProductController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Product created!"})
}

// Update a product
// @Summary Update a product
// @Schemes
// @Description Update a product
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {string} TDB
// @Router /v1/products/:productId [patch]
func (u ProductController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Product updated!"})
}

// Delete a product
// @Summary Delete a product
// @Schemes
// @Description Delete a product
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {string} TDB
// @Router /v1/products/:productId [delete]
func (u ProductController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted!"})
}