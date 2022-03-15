package interfaces

import "github.com/victorschlindwein/gocrud/src/structs"

type InstallmentService interface {
	Create(installment structs.Installment) error
}

type InstallmentRepository interface {
	Create(installment structs.Installment) error
}
