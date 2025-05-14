// Package main - точка входа веб-приложения на Fiber и PostgreSQL.
// Реализует:
//   - Инициализацию всех компонентов системы
//   - Graceful shutdown
//   - Конфигурацию через контекст
package main

import (
	"context"
	"github.com/eeQuillibrium/wobble/pkg/utils"
)

// main - точка входа приложения. Инициализирует:
//   - Глобальный контекст (с возможностью отмены)
//   - Подключение к PostgreSQL
//   - HTTP-сервер на Fiber
//   - Graceful shutdown через utils.Shutdown()
//
// Flow:
//  1. Инициализация БД -> 2. Запуск сервера -> 3. Ожидание shutdown
//
// Сигналы:
//   - SIGINT, SIGTERM: инициируют graceful shutdown
func main() {
	globalCtx, cancel := context.WithCancel(context.Background())

	pool := InitDB(globalCtx)

	srv := InitServer(pool, NewFiberApp(NewFiberTemplates()))

	go srv.Run(globalCtx)

	utils.Shutdown(cancel)
}
