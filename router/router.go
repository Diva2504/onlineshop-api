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
		userRouter.PATCH("/topup")
	}
	categoriesRouter := r.Group("/categories")
  categoriesRouter.Use(middleware.AdminAuth())
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
// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhoZGRoQGpkanNzZC5jb20iLCJleHAiOjE2NTU3NzY2MDcsImlkIjoyfQ.KQC6fNkW5C8lykesYhe4L1KhYpd2Cb7gvKQogR5sgrQ
// }
