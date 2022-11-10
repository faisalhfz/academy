package main

import (
	"ecommerce-api/config"
	"ecommerce-api/handler"
	"ecommerce-api/repository"
	"ecommerce-api/routes"
	"ecommerce-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.AutoMigrate()

	productRepository := repository.NewProductRepository(config.DB)
	cartRepository := repository.NewCartRepository(config.DB)

	productUsecase := usecase.NewProductUsecase(productRepository)
	cartUsecase := usecase.NewCartUsecase(cartRepository)

	productHandler := handler.NewProductHandler(productUsecase, cartUsecase)
	cartHandler := handler.NewCartHandler(cartUsecase)
	transactionHandler := handler.NewTransactionHandler(cartUsecase)

	e := echo.New()
	g := e.Group("/api/v1")

	g.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome to PC Accessories Shop!")
	})

	routes.ProductRoutes(g, productHandler)
	routes.CartRoutes(g, cartHandler)
	routes.TransactionRoutes(g, transactionHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
