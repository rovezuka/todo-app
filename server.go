package todo

import (
	"context"
	"net/http"
	"time"
)

/*
Этот код создает и управляет HTTP-сервером на указанном порту с заданными параметрами, такими как обработчик
(handler), максимальный размер заголовков, время ожидания для чтения и записи запросов. Метод Run запускает
сервер, а метод Shutdown завершает его работу с учетом контекста.
*/

// Server представляет собой структуру для создания и управления HTTP-сервером.
type Server struct {
	httpServer *http.Server
}

// Run запускает HTTP-сервер на указанном порту с указанным обработчиком (handler).
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,       // Устанавливаем адрес и порт, на котором сервер будет слушать запросы.
		Handler:        handler,          // Устанавливаем обработчик (handler) для сервера.
		MaxHeaderBytes: 1 << 20,          // Устанавливаем максимальный размер заголовков запросов (1 MB).
		ReadTimeout:    10 * time.Second, // Устанавливаем время ожидания для чтения запроса (10 секунд).
		WriteTimeout:   10 * time.Second, // Устанавливаем время ожидания для записи ответа (10 секунд).
	}

	return s.httpServer.ListenAndServe() // Запускаем сервер и начинаем слушать указанный порт.
}

// Shutdown завершает работу сервера с учетом контекста (ctx).
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx) // Завершаем работу сервера с учетом указанного контекста.
}
