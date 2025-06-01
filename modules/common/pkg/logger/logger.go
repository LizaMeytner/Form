package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init инициализирует глобальный логгер
func Init() {
	// Настраиваем формат времени
	zerolog.TimeFieldFormat = time.RFC3339

	// Настраиваем вывод в консоль с цветным форматированием
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Устанавливаем глобальный логгер
	log.Logger = zerolog.New(output).With().Timestamp().Logger()
}

// Get возвращает глобальный логгер
func Get() *zerolog.Logger {
	return &log.Logger
}

// WithContext добавляет контекст к логгеру
func WithContext(key, value string) *zerolog.Logger {
	logger := log.With().Str(key, value).Logger()
	return &logger
}

// WithRequestID добавляет ID запроса к логгеру
func WithRequestID(requestID string) *zerolog.Logger {
	return WithContext("request_id", requestID)
}

// WithUserID добавляет ID пользователя к логгеру
func WithUserID(userID string) *zerolog.Logger {
	return WithContext("user_id", userID)
} 