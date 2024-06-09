package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sznborges/projetoToDo/config"
)

var corsConfig = cors.Config{
	AllowOrigins:     "*",
	AllowMethods:     "GET,POST,PUT,DELETE",
	AllowCredentials: true,
	MaxAge:           3600,
}

var staticConfig = fiber.Static{
	Compress:       false,
	ByteRange:      false,
	Browse:         false,
	Download:       true,
	Index:          "layout.csv",
	CacheDuration:  10 * time.Second,
	MaxAge:         604800, // 7 days
	ModifyResponse: nil,
	Next:           nil,
}

func StartHTTP() {
	app := fiber.New(fiber.Config{
		ReadTimeout:           config.GetDuration("HTTP_SERVER_READ_TIMEOUT_MILLIS"),
		WriteTimeout:          config.GetDuration("HTTP_SERVER_WRITE_TIMEOUT_MILLIS"),
		DisableStartupMessage: true,
	})

	app.Use(cors.New(corsConfig))

	app.Static("/", "./public", staticConfig)

	app.Listen(":9000")
}
