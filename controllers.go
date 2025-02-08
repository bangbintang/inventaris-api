package controllers

import (
	"inventaris-api/db"
	"inventaris-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Produk Controllers
func CreateProduct(c *gin.Context) {
    var produk models.Produk
    if err := c.ShouldBindJSON(&produk); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&produk).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": produk})
}

func GetProducts(c *gin.Context) {
    var produk []models.Produk
    if err := db.DB.Find(&produk).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, produk)
}

func GetProductByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var produk models.Produk
    if err := db.DB.First(&produk, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, produk)
}

func GetProductsByCategory(c *gin.Context) {
    kategori := c.Param("kategori")
    var produk []models.Produk
    if err := db.DB.Where("kategori = ?", kategori).Find(&produk).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, produk)
}

func UpdateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var produk models.Produk
    if err := db.DB.First(&produk, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := c.ShouldBindJSON(&produk); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Save(&produk).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": produk})
}

func DeleteProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var produk models.Produk
    if err := db.DB.First(&produk, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := db.DB.Delete(&produk).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// Inventaris Controllers
func GetInventoryByProductID(c *gin.Context) {
    productID, err := strconv.Atoi(c.Param("product_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
        return
    }

    var inventaris models.Inventaris
    if err := db.DB.Where("id_produk = ?", productID).First(&inventaris).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
        return
    }

    c.JSON(http.StatusOK, inventaris)
}

func UpdateInventory(c *gin.Context) {
    productID, err := strconv.Atoi(c.Param("product_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
        return
    }

    var inventaris models.Inventaris
    if err := db.DB.Where("id_produk = ?", productID).First(&inventaris).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
        return
    }

    if err := c.ShouldBindJSON(&inventaris); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Save(&inventaris).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully", "inventory": inventaris})
}

// Pesanan Controllers
func CreateOrder(c *gin.Context) {
    var pesanan models.Pesanan
    if err := c.ShouldBindJSON(&pesanan); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&pesanan).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": pesanan})
}

func GetOrder(c *gin.Context) {
    orderID, err := strconv.Atoi(c.Param("order_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
        return
    }

    var pesanan models.Pesanan
    if err := db.DB.First(&pesanan, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, pesanan)
}
