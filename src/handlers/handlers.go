package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/victorschlindwein/gocrud/src/handlers/installment"
	"github.com/victorschlindwein/gocrud/src/services"
)

func NewHandlerContainer(router fiber.Router, serviceContainer services.ServiceContainer) {
	installment.NewInstallmentHandler(router, serviceContainer).SetRoutes()
}
