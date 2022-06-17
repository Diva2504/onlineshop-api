package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/config"
	"github.com/takadev15/onlineshop-api/controller"
)

func RoutesList() *gin.Engine {
	r := gin.Default()
	db := config.GetDB()
	handler := controller.Handlers{Connect: db}
	userRouter := r.Group("/users")
	{
		userRouter.POST("/login", handler.UserLogin)
		userRouter.POST("/register", handler.UserRegister)
		userRouter.PATCH("/topup")
	}
	categoriesRouter := r.Group("/categories")
	{
		categoriesRouter.GET("/", handler.GetCategory)
		categoriesRouter.POST("/", handler.CreateCategory)
		categoriesRouter.PATCH("/:id")
		categoriesRouter.DELETE("/:id")
	}

	productRouter := r.Group("/products")
	{
		productRouter.GET("/", handler.GetAllProduct)
		productRouter.GET("/:id", handler.GetProduct)
		productRouter.POST("/", handler.CreateProduct)
		productRouter.PUT("/:id", handler.UpdateProduct)
		productRouter.DELETE("/:id", handler.DeleteProduct)
	}

	transactionRouter := r.Group("/transactions")
	{
		transactionRouter.POST("/", handler.CreateTransaction)
		transactionRouter.GET("/my-transaction", handler.GetforUser)
		// transactionRouter.GET("/:user_id", handler.GetforAdmin)
	}
	return r
}

// {
//     "full_name": "Dagga",
//     "email": "hhddh@jdjd.com",
//     "password": "uuuuuu"
// }
