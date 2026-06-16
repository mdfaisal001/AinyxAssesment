package handler

import (
	"strconv"

	"go-user-api/internal/logger"
	"go-user-api/internal/models"
	"go-user-api/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var validate = validator.New()

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(
	service *service.UserService,
) *UserHandler {

	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) CreateUser(
	c *fiber.Ctx,
) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"Creating user",
		zap.String("name", req.Name),
		zap.String("dob", req.DOB),
	)

	user, err := h.Service.CreateUser(
		req.Name,
		req.DOB,
	)

	if err != nil {
		logger.Log.Error(
			"Failed to create user",
			zap.Error(err),
		)

		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"User created successfully",
		zap.Int32("user_id", user.ID),
		zap.String("name", user.Name),
	)

	return c.Status(201).JSON(user)
}
func (h *UserHandler) GetAllUsers(
	c *fiber.Ctx,
) error {

	logger.Log.Info("Fetching all users")

	users, err := h.Service.GetAllUsers()

	if err != nil {

		logger.Log.Error(
			"Failed to fetch users",
			zap.Error(err),
		)

		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	var response []models.UserResponse

	for _, user := range users {
		response = append(response, models.UserResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Format("2006-01-02"),
			Age:  service.CalculateAge(user.Dob),
		})
	}

	logger.Log.Info(
		"Users fetched successfully",
		zap.Int("count", len(response)),
	)

	return c.JSON(response)
}
func (h *UserHandler) GetUserByID(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	logger.Log.Info(
		"Fetching user by id",
		zap.Int("user_id", id),
	)

	user, err := h.Service.GetUserByID(int32(id))

	if err != nil {

		logger.Log.Error(
			"User not found",
			zap.Int("user_id", id),
			zap.Error(err),
		)

		return c.Status(404).JSON(
			fiber.Map{
				"error": "user not found",
			},
		)
	}

	response := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  service.CalculateAge(user.Dob),
	}

	logger.Log.Info(
		"User fetched successfully",
		zap.Int32("user_id", user.ID),
	)

	return c.JSON(response)
}
func (h *UserHandler) DeleteUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	logger.Log.Info(
		"Deleting user",
		zap.Int("user_id", id),
	)

	err = h.Service.DeleteUser(
		int32(id),
	)

	if err != nil {

		logger.Log.Error(
			"Failed to delete user",
			zap.Int("user_id", id),
			zap.Error(err),
		)

		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"User deleted successfully",
		zap.Int("user_id", id),
	)

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) UpdateUser(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"Updating user",
		zap.Int("user_id", id),
		zap.String("name", req.Name),
	)

	user, err := h.Service.UpdateUser(
		int32(id),
		req.Name,
		req.DOB,
	)

	if err != nil {

		logger.Log.Error(
			"Failed to update user",
			zap.Int("user_id", id),
			zap.Error(err),
		)

		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	logger.Log.Info(
		"User updated successfully",
		zap.Int32("user_id", user.ID),
		zap.String("name", user.Name),
	)

	return c.JSON(user)
}