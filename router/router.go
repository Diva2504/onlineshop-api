package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/config"
	"github.com/takadev15/onlineshop-api/controller"
	"github.com/takadev15/onlineshop-api/middleware"
)

func RoutesList() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	handler := controller.Handlers{Connect: db}
	userRouter := r.Group("/users")
	{
		userRouter.POST("/login", handler.UserLogin)
		userRouter.POST("/register", handler.UserRegister)
		userRouter.PATCH("/topup", middleware.Authentication(), handler.UserTopup)
	}
	categoriesRouter := r.Group("/categories")
  categoriesRouter.Use(middleware.Authentication(), middleware.AdminAuth())
	{
		categoriesRouter.GET("/", handler.GetCategory)
		categoriesRouter.POST("/", handler.CreateCategory)
		categoriesRouter.PATCH("/:id", handler.UpdateCategory)
		categoriesRouter.DELETE("/:id", handler.DeleteCategory)
	}

	productRouter := r.Group("/products")
  productRouter.Use(middleware.Authentication(), middleware.AdminAuth())
	{
		productRouter.GET("/", handler.GetAllProduct)
		productRouter.GET("/:id", handler.GetProduct)
		productRouter.POST("/", handler.CreateProduct)
		productRouter.PUT("/:id", handler.UpdateProduct)
		productRouter.DELETE("/:id", handler.DeleteProduct)
	}

	transactionRouter := r.Group("/transactions")
  transactionRouter.Use(middleware.Authentication())
	{
		transactionRouter.POST("/", handler.CreateTransaction)
		transactionRouter.GET("/my-transaction", handler.GetforUser)
	}
	return r
}

