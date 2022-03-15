package repositories

import (
	"github.com/victorschlindwein/gocrud/src/interfaces"
	"github.com/victorschlindwein/gocrud/src/repositories/installment"
	"gorm.io/gorm"
)

type RepositoryContainer struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func GetRepositories(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		InstallmentRepository: installment.NewInstallmentRepository(db),
	}
}
