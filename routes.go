package routes

import (
	"inventaris-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    // Produk Routes
    r.POST("/products", controllers.CreateProduct)
    r.GET("/products", controllers.GetProducts)
    r.GET("/products/:id", controllers.GetProductByID)
    r.GET("/products/category/:kategori", controllers.GetProductsByCategory)
    r.PUT("/products/:id", controllers.UpdateProduct)
    r.DELETE("/products/:id", controllers.DeleteProduct)

    // Inventaris Routes
    r.GET("/inventory/:product_id", controllers.GetInventoryByProductID)
    r.PUT("/inventory/:product_id", controllers.UpdateInventory)

    // Pesanan Routes
    r.POST("/orders", controllers.CreateOrder)
    r.GET("/orders/:order_id", controllers.GetOrder)
}
