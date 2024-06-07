package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(app *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, os.Interrupt)
		<-sigInt

		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v\n", err)
		}

		close(idleConnsClosed)
	}()

	fiberConnUrl, _ := ConnectionUrlBuilder("fiber")

	if err := app.Listen(fiberConnUrl); err != nil {
		log.Printf("Oops... Server is not starting! Reason: %v\n", err)
	}

	<-idleConnsClosed
}

func StartServer(app *fiber.App) {
	fiberConnUrl, _ := ConnectionUrlBuilder("fiber")

	if err := app.Listen(fiberConnUrl); err != nil {
		log.Printf("Oops... Server is not starting! Reason: %v\n", err)
	}
}
