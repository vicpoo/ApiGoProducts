package main

import (
	"demob/src/application"
	"demob/src/domain"
	"demob/src/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos
	db, err := infrastructure.ConnectDB()
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
	defer db.Close()

	// Inicializar repositorio
	repo := infrastructure.NewProductRepository(db)

	// Inicializar Gin
	r := gin.Default()

	// Crear casos de uso
	createProduct := application.CreateProductUseCase{Repo: repo}
	viewAllProducts := application.ViewAllProductsUseCase{Repo: repo}
	updateProduct := application.UpdateProductUseCase{Repo: repo}
	deleteProduct := application.DeleteProductUseCase{Repo: repo}

	// Rutas de la API
	r.POST("/products", func(c *gin.Context) {
		var product domain.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := createProduct.Execute(product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Producto creado con éxito"})
	})

	r.GET("/products", func(c *gin.Context) {
		products, err := viewAllProducts.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		var product domain.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		product.ID = id
		if err := updateProduct.Execute(product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado con éxito"})
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}
		if err := deleteProduct.Execute(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado con éxito"})
	})

	// Correr el servidor en el puerto 8080
	r.Run(":8000") // Cambia el puerto si es necesario
}
