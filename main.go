package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andy-gate/artaka-tenant/controllers"
	"github.com/andy-gate/artaka-tenant/middlewares"
	"github.com/andy-gate/artaka-tenant/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	fmt.Printf("Started at : %3v \n", time.Now())

	if err := godotenv.Load(`.env`); err != nil {
		panic(err)
	}

	models.InitGormPostgres()
	defer models.MPosGORM.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api_tenant")
	api.POST("/login", controllers.Login)

	protected:= router.Group("/api_tenant/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/tenant_list", controllers.TenantList)
	protected.POST("/tenant_list_dropdown", controllers.ActiveTenantList)
	protected.POST("/dashboard", controllers.Dashboard)
	protected.POST("/product_list", controllers.ProductList)
	protected.POST("/sales_list", controllers.SalesList)
	protected.POST("/sales_list_detail", controllers.SalesListDetail)
	protected.POST("/change_status", controllers.ChangeTenantRefCode)
	protected.POST("/sales_list_realtime", controllers.SalesListRealTime)
	protected.POST("/sales_list_detail_realtime", controllers.SalesListDetailRealTime)
	protected.POST("/product_list_realtime", controllers.ProductListRealTime)
	
	fmt.Printf("Listening to port %s", os.Getenv("PORT1"))
	router.Run(":" + os.Getenv("PORT1"))
}