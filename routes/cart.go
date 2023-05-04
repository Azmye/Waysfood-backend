package routes

import (
	"waysfood/handlers"
	"waysfood/pkg/mysql"
	"waysfood/repositories"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(userRepository)

	e.GET("/carts", h.FindCarts)
	e.GET("/cart/:id", h.GetCart)
	e.DELETE("/cart/:id", h.DeleteCart)
}
