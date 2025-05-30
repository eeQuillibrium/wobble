// Package http предоставляет HTTP-интерфейс для взаимодействия с API сервиса.
// Включает роуты для работы с аккаунтами, корзиной, магазином и главной страницей.
// Все обработчики реализованы как отдельные подпакеты для соблюдения принципа разделения ответственности.
package http

import (
	"github.com/eeQuillibrium/wobble/internal/controller"
	"github.com/eeQuillibrium/wobble/internal/transport/http/account"
	"github.com/eeQuillibrium/wobble/internal/transport/http/contact"
	"github.com/eeQuillibrium/wobble/internal/transport/http/store"
)

// API объединяет все HTTP-обработчики сервиса.
// Содержит интерфейсы всех сервисов:
//   - Account: управление пользователями
//   - Cart: операции с корзиной
//   - Store: товары и категории
//   - Index: главная страница
type API struct {
	account account.IAccountAPI
	store   store.IStoreAPI
	contact contact.IContactAPI
}

// New создает инстанс API с инициализированными обработчиками.
// Принимает:
//   - ctrl: основной контроллер приложения, содержащий зависимости для всех подсистем.
//
// Возвращает готовую структуру API, которую можно использовать для регистрации роутов.
// Пример:
//
//	api := http.New(ctrl)
//	router.Register("/account", api.account.Handler())
func New(ctrl controller.Controller) API {
	return API{
		account: account.New(ctrl.Account),
		store:   store.New(ctrl.Store),
		contact: contact.New(ctrl.Contact),
	}
}
