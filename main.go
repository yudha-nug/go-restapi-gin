package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yudha-nug/go-restapi-gin/controllers/productcontroller"
	"github.com/yudha-nug/go-restapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("api/products", productcontroller.Index)
	r.GET("api/product/:id", productcontroller.Show)
	r.POST("api/product", productcontroller.Create)
	r.PUT("api/product/:id", productcontroller.Update)
	r.DELETE("api/product", productcontroller.Delete)

	r.Run()
}
