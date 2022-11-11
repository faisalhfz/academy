package handler

import (
	"ecommerce-api/entity"
	"ecommerce-api/entity/request"
	"ecommerce-api/entity/response"
	"ecommerce-api/usecase"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	pUsecase *usecase.ProductUsecase
	cUsecase *usecase.CartUsecase
}

func NewProductHandler(pUsecase *usecase.ProductUsecase, cUsecase *usecase.CartUsecase) *ProductHandler {
	return &ProductHandler{pUsecase: pUsecase, cUsecase: cUsecase}
}

func createProductRequest(ctx echo.Context) (request.CreateProductRequest, error) {
	name := ctx.FormValue("name")
	image, err := ctx.FormFile("image")
	if err != nil {
		return request.CreateProductRequest{}, err
	}
	src, err := image.Open()
	if err != nil {
		return request.CreateProductRequest{}, err
	}
	defer src.Close()
	dst, err := os.Create(image.Filename)
	if err != nil {
		return request.CreateProductRequest{}, err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return request.CreateProductRequest{}, err
	}
	absPath, _ := filepath.Abs("./")
	imgPath := absPath + "\\" + image.Filename
	price, err := strconv.Atoi(ctx.FormValue("price"))
	if err != nil {
		return request.CreateProductRequest{}, err
	}
	category := ctx.FormValue("category")
	description := ctx.FormValue("description")
	return request.CreateProductRequest{Name: name, Image: imgPath, Price: price, Category: category, Description: description}, nil
}

func (pHandler ProductHandler) PostProductHandler(ctx echo.Context) error {
	productRequest, err := createProductRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
	}
	if err := pHandler.pUsecase.CreateProduct(productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to create product",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Product created successfully",
		Data:    nil,
	})
}

func (pHandler ProductHandler) GetProductsHandler(ctx echo.Context) error {
	category := ctx.FormValue("category")
	minPrice, _ := strconv.Atoi(ctx.FormValue("minprice"))
	maxPrice, _ := strconv.Atoi(ctx.FormValue("maxprice"))
	filter := request.ProductFilterRequest{Category: category, MinPrice: minPrice, MaxPrice: maxPrice}
	products, err := pHandler.pUsecase.GetProducts(filter)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to get products",
			Data:    nil,
		})
	}
	if len(products) == 0 {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "No products found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get all products",
		Data:    products,
	})
}

func (pHandler ProductHandler) GetProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	product, err := pHandler.pUsecase.GetProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}
	productResponse := response.GetProductDetailResponse{}
	copier.Copy(&productResponse, product)
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get product",
		Data:    productResponse,
	})
}

func (pHandler ProductHandler) PutProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	product, err := pHandler.pUsecase.GetProductById(id)
	_ = product
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}
	productRequest, err := createProductRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
	}
	if err := pHandler.pUsecase.EditProductById(id, productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to update product",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully updated product",
		Data:    nil,
	})
}

func (pHandler ProductHandler) DeleteProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	product, err := pHandler.pUsecase.GetProductById(id)
	_ = product
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}
	if err := pHandler.pUsecase.DeleteProductById(id); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to delete product",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully deleted product",
		Data:    nil,
	})
}

func (pHandler ProductHandler) PostProductToCartHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	product, err := pHandler.pUsecase.GetProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}
	productEntry := entity.Product{}
	copier.Copy(&productEntry, &product)
	cart, err := pHandler.cUsecase.GetCart()
	_ = cart
	if err != nil {
		newCart := request.CreateCartRequest{Products: append([]entity.Product{}, productEntry), TotalPrice: productEntry.Price, IsCheckout: false}
		if err := pHandler.cUsecase.CreateCart(newCart); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Failed to create cart",
				Data:    nil,
			})
		}
		return ctx.JSON(http.StatusCreated, response.BaseResponse{
			Code:    http.StatusCreated,
			Message: "Successfully created cart and added product to cart",
			Data:    nil,
		})
	}
	if err := pHandler.cUsecase.AddProductToCart(productEntry); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to add product to cart",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Successfully added product to cart",
		Data:    nil,
	})
}
