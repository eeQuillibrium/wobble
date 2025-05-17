package account

import (
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/pkg/logger"
	"github.com/eeQuillibrium/wobble/pkg/utils"
	"github.com/gofiber/fiber/v3"
)

func (a *API) Register(c fiber.Ctx) error {
	var reg dto.Register
	if err := c.Bind().JSON(&reg); err != nil {
		logger.Ctx(c.Context()).Warn("failed to bind dto.Register")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Все поля обязательны для заполнения",
		})
	}

	user, err := a.ctrl.Register(c.Context(), reg)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка при регистрации",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Пользователь успешно зарегистрирован",
		"user_id": user.ID,
	})
}

func (a *API) Auth(c fiber.Ctx) error {
	var log dto.Login
	if err := c.Bind().JSON(&log); err != nil {
		logger.Ctx(c.Context()).Warn("failed to bind dto.Login")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка при парсинге json-a",
		})
	}

	token, userID, err := a.ctrl.Auth(c.Context(), log)
	if err != nil {
		logger.Ctx(c.Context()).Warn("failed to login")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ошибка при авторизации",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		Secure:   true,
		MaxAge:   86400,
	})

	return c.JSON(fiber.Map{
		"token":   token,
		"user_id": userID,
	})
}

func (a *API) Logout(c fiber.Ctx) error {
	utils.ClearCookies(c, "jwt")

	return c.Redirect().To("/")
}
