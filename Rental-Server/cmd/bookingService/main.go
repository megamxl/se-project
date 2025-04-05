package main

import (
	"github.com/megamxl/se-project/Rental-Server/internal/config"
	"github.com/megamxl/se-project/Rental-Server/internal/middleware"
	"os"
)

func main() {

	h := config.BasicServerSetup()

	hWithMiddleware := middleware.BookingsServiceMiddleware(h)

	config.ListenAndServeServer(hWithMiddleware, os.Getenv("WEB_HOST")+":"+os.Getenv("WEB_PORT"))
}
