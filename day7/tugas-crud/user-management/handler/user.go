package handler

import (
	"user-management/entity/response"
	"user-management/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usercase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: usercase}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to create user",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "User created successfully",
		Data:    nil,
	})
}

func (h UserHandler) GetList(ctx *fiber.Ctx) error {
	qname := ctx.Query("name")
	users, err := h.userUsecase.GetList(qname)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to get users",
			Data:    nil,
		})
	}
	if len(users) <= 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Code:    fiber.StatusNotFound,
			Message: "No user found",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Successfully get all users",
		Data:    users,
	})
}

func (h UserHandler) GetUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid id",
			Data:    nil,
		})
	}
	user, err := h.userUsecase.GetUser(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Code:    fiber.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Successfully found user",
		Data:    user,
	})
}

func (h UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid id",
			Data:    nil,
		})
	}
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.UpdateUser(id, userRequest); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Code:    fiber.StatusNotFound,
			Message: "Failed to update user. User not found",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Successfully updated user",
		Data:    nil,
	})
}

func (h UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid id",
			Data:    nil,
		})
	}
	if err := h.userUsecase.DeleteUser(id); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.BaseResponse{
			Code:    fiber.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Successfully deleted user",
		Data:    nil,
	})
}
