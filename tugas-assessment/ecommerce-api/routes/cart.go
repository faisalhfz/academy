package routes

import (
	"ecommerce-api/handler"

	"github.com/labstack/echo/v4"
)

func CartRoutes(g *echo.Group, cartHandler *handler.CartHandler) {
	g.GET("/cart", cartHandler.GetCartHandler)
	g.POST("/cart", cartHandler.PostCartHandler)
	g.DELETE("/cart", cartHandler.DeleteCartHandler)
	// g.DELETE("/cart/:id", cartHandler.DeleteProductFromCartHandler)
}
