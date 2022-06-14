package router

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/onlineshop-api/config"
	"github.com/takadev15/onlineshop-api/controller"
	"github.com/takadev15/onlineshop-api/middleware"
)

func RoutesList() *gin.Engine{
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
  categoriesRouter.Use(middleware.Authentication(), middleware.AdminAuth())
  {
    categoriesRouter.GET("/")
    categoriesRouter.POST("/")
    categoriesRouter.PATCH("/:id")
    categoriesRouter.DELETE("/:id")
  }
  return r
}

// {
//     "full_name": "Dagga",
//     "email": "hhddh@jdjd.com",
//     "password": "uuuuuu" 
// }
