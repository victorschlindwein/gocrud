package services

import (
	"github.com/victorschlindwein/gocrud/src/interfaces"
	"github.com/victorschlindwein/gocrud/src/repositories"
	"github.com/victorschlindwein/gocrud/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(RepositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		InstallmentService: installment.NewInstallmentService(RepositoryContainer),
	}
}
