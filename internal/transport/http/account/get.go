package account

import (
	"github.com/gofiber/fiber/v3"
	"time"
)

func (a *API) GetUser(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint64)
	role := c.Locals("role").(string)

	user, err := a.ctrl.GetUserByID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось загрузить данные пользователя",
		})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"login": user.Login,
		"role":  role,
	})
}

func (a *API) GetOrders(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint64)

	orders, err := a.ctrl.GetOrdersByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось загрузить историю заказов",
		})
	}

	// Преобразуем заказы в нужный формат
	var response []fiber.Map
	for _, order := range orders {
		response = append(response, fiber.Map{
			"id":              order.ID,
			"orderDate":       order.CreatedAt.Format(time.RFC3339),
			"status":          order.Status,
			"totalAmount":     order.TotalAmount,
			"items":           order.Items,
			"deliveryAddress": order.DeliveryAddress,
		})
	}

	return c.JSON(response)
}
